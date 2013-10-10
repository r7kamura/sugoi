package sugoi

import (
	"encoding/xml"
)

type Title struct {
	Cat string
	Comment string
	FirstCh string
	FirstEndMonth string
	FirstEndYear string
	FirstMonth string
	FirstYear string
	Keywords string
	LastUpdate string
	ShortTitle string
	SubTitles string
	TID string
	Title string
	TitleEN string
	TitleFlag string
	TitleYomi string
	UserPoint string
	UserPointRank string
}

func NewTitles(data []byte) ([]*Title, error) {
	return NewTitlesParser(data).Parse()
}

type TitlesParser struct {
	Data []byte
}

func NewTitlesParser(data []byte) *TitlesParser {
	return &TitlesParser{Data: data}
}

func (parser *TitlesParser) Parse() ([]*Title, error) {
	response, err := parser.TitleLookupResponse()
	if err != nil {
		return nil, err
	}
	titles := make([]*Title, len(response.TitleItems.TitleItem))
	for i, titleItem := range response.TitleItems.TitleItem {
		titles[i] = &Title{
			Cat:           titleItem.Cat,
			Comment:       titleItem.Comment,
			FirstCh:       titleItem.FirstCh,
			FirstEndMonth: titleItem.FirstEndMonth,
			FirstEndYear:  titleItem.FirstEndYear,
			FirstMonth:    titleItem.FirstMonth,
			FirstYear:     titleItem.FirstYear,
			Keywords:      titleItem.Keywords,
			LastUpdate:    titleItem.LastUpdate,
			ShortTitle:    titleItem.ShortTitle,
			SubTitles:     titleItem.SubTitles,
			TID:           titleItem.TID,
			Title:         titleItem.Title,
			TitleEN:       titleItem.TitleEN,
			TitleFlag:     titleItem.TitleFlag,
			TitleYomi:     titleItem.TitleYomi,
			UserPoint:     titleItem.UserPoint,
			UserPointRank: titleItem.UserPointRank,
		}
	}
	return titles, nil
}

func (parser *TitlesParser) TitleLookupResponse() (*titleLookupResponse, error) {
	var response titleLookupResponse
	return &response, xml.Unmarshal(parser.Data, &response)
}

type titleLookupResponse struct {
	TitleItems titleItems `xml:"TitleItems"`
}

type titleItems struct {
	TitleItem []titleItem `xml:"TitleItem"`
}

type titleItem struct {
	Cat string
	Comment string
	FirstCh string
	FirstEndMonth string
	FirstEndYear string
	FirstMonth string
	FirstYear string
	Keywords string
	LastUpdate string
	ShortTitle string
	SubTitles string
	TID string
	Title string
	TitleEN string
	TitleFlag string
	TitleYomi string
	UserPoint string
	UserPointRank string
}
