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

func NewTitle(data []byte) (title *Title, err error) {
	return NewTitleParser(data).Parse()
}

type TitleParser struct {
	Data []byte
}

func NewTitleParser(data []byte) *TitleParser {
	return &TitleParser{Data: data}
}

func (parser *TitleParser) Parse() (title *Title, err error) {
	response, err := parser.TitleLookupResponse()
	title = &Title{
		Cat:           response.TitleItems.TitleItem.Cat,
		Comment:       response.TitleItems.TitleItem.Comment,
		FirstCh:       response.TitleItems.TitleItem.FirstCh,
		FirstEndMonth: response.TitleItems.TitleItem.FirstEndMonth,
		FirstEndYear:  response.TitleItems.TitleItem.FirstEndYear,
		FirstMonth:    response.TitleItems.TitleItem.FirstMonth,
		FirstYear:     response.TitleItems.TitleItem.FirstYear,
		Keywords:      response.TitleItems.TitleItem.Keywords,
		LastUpdate:    response.TitleItems.TitleItem.LastUpdate,
		ShortTitle:    response.TitleItems.TitleItem.ShortTitle,
		SubTitles:     response.TitleItems.TitleItem.SubTitles,
		TID:           response.TitleItems.TitleItem.TID,
		Title:         response.TitleItems.TitleItem.Title,
		TitleEN:       response.TitleItems.TitleItem.TitleEN,
		TitleFlag:     response.TitleItems.TitleItem.TitleFlag,
		TitleYomi:     response.TitleItems.TitleItem.TitleYomi,
		UserPoint:     response.TitleItems.TitleItem.UserPoint,
		UserPointRank: response.TitleItems.TitleItem.UserPointRank,
	}
	return
}

func (parser *TitleParser) TitleLookupResponse() (*titleLookupResponse, error) {
	var response titleLookupResponse
	return &response, xml.Unmarshal(parser.Data, &response)
}

type titleLookupResponse struct {
	TitleItems titleItems `xml:"TitleItems"`
}

type titleItems struct {
	TitleItem titleItem `xml:"TitleItem"`
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
