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
	table := map[string]string {
		"Command": "ProgLookup",
		"Range": playedRange,
		"PID": builder.ID(),
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
