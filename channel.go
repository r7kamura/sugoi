package sugoi

import "encoding/xml"

type Channel struct {
	ChannelGroupID string `xml:"ChGID"`
	Comment        string `xml:"ChComment"`
	EPGURL         string `xml:"ChEPGURL"`
	ID             string `xml:"ChID"`
	IEPGName       string `xml:"ChiEPGName"`
	Name           string `xml:"ChName"`
	Number         string `xml:"ChNumber"`
	URL            string `xml:"ChURL"`
	UpdatedAt      string `xml:"LastUpdate"`
}

func NewChannels(data []byte) ([]*Channel, error) {
	return NewChannelsParser(data).Parse()
}

type ChannelsParser struct {
	Data []byte
}

func NewChannelsParser(data []byte) *ChannelsParser {
	return &ChannelsParser{Data: data}
}

func (parser *ChannelsParser) Parse() ([]*Channel, error) {
	response, err := parser.ChannelLookupResponse()
	if err != nil {
		return nil, err
	}
	channels := make([]*Channel, len(response.ChannelItems.Channels))
	for i, channel := range response.ChannelItems.Channels {
		channels[i] = &channel
	}
	return channels, err
}

func (parser *ChannelsParser) ChannelLookupResponse() (*channelLookupResponse, error) {
	var response channelLookupResponse
	return &response, xml.Unmarshal(parser.Data, &response)
}

type channelLookupResponse struct {
	ChannelItems channelItems `xml:"ChItems"`
}

type channelItems struct {
	Channels []Channel `xml:"ChItem"`
}
