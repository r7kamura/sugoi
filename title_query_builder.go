package sugoi

import "net/url"

type TitleQueryBuilder struct {
	Options map[string]string
}

func NewTitleQueryBuilder(pairs ...string) *TitleQueryBuilder {
	return &TitleQueryBuilder{Options: convertKeyValuePairsToHash(pairs...)}
}

func (builder *TitleQueryBuilder) Build() (string, error) {
	lastUpdate, err := builder.LastUpdate()
	if err != nil {
		return "", err
	}
	table := map[string]string {
		"Command": "TitleLookup",
		"LastUpdate": lastUpdate,
		"TID": builder.TID(),
	}
	values := url.Values{}
	for key, value := range table {
		if value != "" {
			values.Set(key, value)
		}
	}
	return values.Encode(), nil
}

func (builder *TitleQueryBuilder) TID() string {
	if id := builder.Options["id"]; id != "" {
		return id
	} else {
		return "*"
	}
}

func (builder *TitleQueryBuilder) LastUpdate() (string, error) {
	from, err := builder.UpdatedFrom()
	if err != nil {
		return "", nil
	}
	to, err := builder.UpdatedTo()
	if err != nil {
		return "", nil
	}
	if from != "" || to != "" {
		return from + "-" + to, nil
	} else {
		return "", nil
	}
}

func (builder *TitleQueryBuilder) UpdatedFrom() (string, error) {
	return builder.FormattedTimeOf("updatedFrom")
}

func (builder *TitleQueryBuilder) UpdatedTo() (string, error) {
	return builder.FormattedTimeOf("updatedTo")
}

func (builder *TitleQueryBuilder) FormattedTimeOf(key string) (string, error) {
	if str, ok := builder.Options[key]; ok {
		return convertRFC3339ToSyoboiFormat(str)
	} else {
		return "", nil
	}
}
