package sugoi

import "net/url"

type ProgramQueryBuilder struct {
	Options map[string]string
}

func NewProgramQueryBuilder(pairs ...string) *ProgramQueryBuilder {
	options := map[string]string{}
	for i := 0; i + 1 < len(pairs); i += 2 {
		options[pairs[i]] = pairs[i + 1]
	}
	return &ProgramQueryBuilder{Options: options}
}

func (builder *ProgramQueryBuilder) Build() (string, error) {
	playedRange, err := builder.PlayedRange()
	if err != nil {
		return "", err
	}
	stTime, err := builder.StTime()
	if err != nil {
		return "", err
	}
	table := map[string]string {
		"ChID": builder.ChID(),
		"Command": "ProgLookup",
		"PID": builder.ID(),
		"Range": playedRange,
		"StTime": stTime,
	}
	values := url.Values{}
	for key, value := range table {
		if value != "" {
			values.Set(key, value)
		}
	}
	return values.Encode(), nil
}

func (builder *ProgramQueryBuilder) ID() string {
	return builder.Options["id"]
}

func (builder *ProgramQueryBuilder) ChID() string {
	return builder.Options["channelID"]
}

func (builder *ProgramQueryBuilder) StTime() (string, error) {
	to, err := builder.StartedTo()
	if err != nil {
		return "", err
	}
	if to == "" {
		return "", nil
	} else {
		return "-" + to, nil
	}
}

func (builder *ProgramQueryBuilder) StartedTo() (string, error) {
	return builder.FormattedTimeOf("startedTo")
}

func (builder *ProgramQueryBuilder) PlayedRange() (string, error) {
	from, err := builder.PlayedFrom()
	if err != nil {
		return "", err
	}
	to, err := builder.PlayedTo()
	if err != nil {
		return "", err
	}
	if from != "" || to != "" {
		return from + "-" + to, nil
	} else {
		return "", nil
	}
}

func (builder *ProgramQueryBuilder) PlayedFrom() (string, error) {
	return builder.FormattedTimeOf("playedFrom")
}

func (builder *ProgramQueryBuilder) PlayedTo() (string, error) {
	return builder.FormattedTimeOf("playedTo")
}

func (builder *ProgramQueryBuilder) FormattedTimeOf(key string) (string, error) {
	if str, ok := builder.Options[key]; ok {
		return convertRFC3339ToSyoboiFormat(str)
	} else {
		return "", nil
	}
}
