package sugoi

import "encoding/xml"

type Program struct {
	ChannelID      string `xml:"ChID"`
	Comment        string `xml:"ProgItem"`
	Count          string `xml:"Count"`
	Deleted        string `xml:"Deleted"`
	EditedAt       string `xml:"EdTime"`
	Flag           string `xml:"Flag"`
	ID             string `xml:"PID"`
	Item           string `xml:"ProgComment"`
	JoinedSubTitle string `xml:"STSubTitle"`
	Revision       string `xml:"Revision"`
	StartedAt      string `xml:"StTime"`
	StartedOffset  string `xml:"StOffset"`
	SubTitle       string `xml:"SubTitle"`
	TitleID        string `xml:"TID"`
	UpdatedAt      string `xml:"LastUpdate"`
	Warning        string `xml:"Warn"`
}

func NewPrograms(data []byte) ([]*Program, error) {
	return NewProgramsParser(data).Parse()
}

type ProgramsParser struct {
	Data []byte
}

func NewProgramsParser(data []byte) *ProgramsParser {
	return &ProgramsParser{Data: data}
}

func (parser *ProgramsParser) Parse() ([]*Program, error) {
	response, err := parser.ProgramLookupResponse()
	if err != nil {
		return nil, err
	}
	programs := make([]*Program, len(response.ProgramItems.Programs))
	for i, program := range response.ProgramItems.Programs {
		programs[i] = &program
	}
	return programs, err
}

func (parser *ProgramsParser) ProgramLookupResponse() (*programLookupResponse, error) {
	var response programLookupResponse
	return &response, xml.Unmarshal(parser.Data, &response)
}

type programLookupResponse struct {
	ProgramItems programItems `xml:"ProgItems"`
}

type programItems struct {
	Programs []Program `xml:"ProgItem"`
}
