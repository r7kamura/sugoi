package sugoi

import "net/url"

type ProgramQueryBuilder struct {
	Options map[string]string
}

func NewProgramQueryBuilder(pairs ...string) *ProgramQueryBuilder {
	return &ProgramQueryBuilder{Options: convertKeyValuePairsToHash(pairs...)}
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
		"Count": builder.Count(),
		"JOIN": builder.Join(),
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

func (builder *ProgramQueryBuilder) Count() string {
	return builder.Options["count"]
}

func (builder *ProgramQueryBuilder) Join() string {
	if builder.Options["join"] == "0" {
		return ""
	} else {
		return "SubTitles"
	}
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
