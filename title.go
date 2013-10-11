package sugoi

import "encoding/xml"

type Title struct {
	CategoryID    string `xml:"Cat"`
	Comment       string `xml:"Comment"`
	FirstChannel  string `xml:"FirstCh"`
	FirstEndMonth string `xml:"FirstEndMonth"`
	FirstEndYear  string `xml:"FirstEndYear"`
	FirstMonth    string `xml:"FirstMonth"`
	FirstYear     string `xml:"FirstYear"`
	ID            string `xml:"TID"`
	Keywords      string `xml:"Keywords"`
	ShortTitle    string `xml:"ShortTitle"`
	SubTitles     string `xml:"SubTitles"`
	Title         string `xml:"Title"`
	TitleEnglish  string `xml:"TitleEN"`
	TitleFlag     string `xml:"TitleFlag"`
	TitleYomi     string `xml:"TitleYomi"`
	UpdatedAt     string `xml:"LastUpdate"`
	UserPoint     string `xml:"UserPoint"`
	UserPointRank string `xml:"UserPointRank"`
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
	titles := make([]*Title, len(response.TitleItems.Titles))
	for i, title := range response.TitleItems.Titles {
		titles[i] = &title
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
	Titles []Title `xml:"TitleItem"`
}
