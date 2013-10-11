package sugoi

import "net/url"

type ChannelQueryBuilder struct {
	Options map[string]string
}

func NewChannelQueryBuilder(pairs ...string) *ChannelQueryBuilder {
	return &ChannelQueryBuilder{Options: convertKeyValuePairsToHash(pairs...)}
}

func (builder *ChannelQueryBuilder) Build() (string, error) {
	lastUpdate, err := builder.LastUpdate()
	if err != nil {
		return "", err
	}
	table := map[string]string {
		"ChID": builder.ChID(),
		"Command": "ChLookup",
		"LastUpdate": lastUpdate,
	}
	values := url.Values{}
	for key, value := range table {
		if value != "" {
			values.Set(key, value)
		}
	}
	return values.Encode(), nil
}

func (builder *ChannelQueryBuilder) ChID() string {
	return builder.Options["id"]
}

func (builder *ChannelQueryBuilder) LastUpdate() (string, error) {
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

func (builder *ChannelQueryBuilder) UpdatedFrom() (string, error) {
	return builder.FormattedTimeOf("updatedFrom")
}

func (builder *ChannelQueryBuilder) UpdatedTo() (string, error) {
	return builder.FormattedTimeOf("updatedTo")
}

func (builder *ChannelQueryBuilder) FormattedTimeOf(key string) (string, error) {
	if str, ok := builder.Options[key]; ok {
		return convertRFC3339ToSyoboiFormat(str)
	} else {
		return "", nil
	}
}
