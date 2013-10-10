package sugoi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var defaultBaseURL string = "http://cal.syoboi.jp"

type Client struct {
	BaseURL string
}

func NewClient() *Client {
	return &Client{BaseURL: defaultBaseURL}
}

func NewClientWithBaseURL(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (client *Client) GetTitleByID(id string) (*Title, error) {
	titles, err := client.GetTitles("/db.php", "TID", id, "Command", "TitleLookup")
	if err != nil {
		return nil, err
	}
	if len(titles) == 0 {
		return nil, &NotFoundError{}
	}
	return titles[0], nil
}

func (client *Client) GetTitlesUpdatedIn(from, to time.Time) ([]*Title, error) {
	return client.GetTitles(
		"/db.php",
		"Command",
		"TitleLookup",
		"TID",
		"*",
		"LastUpdate",
		fmt.Sprintf(
			"%04d%02d%02d_%02d%02d%02d-%04d%02d%02d_%02d%02d%02d",
			from.Year(),
			from.Month(),
			from.Day(),
			from.Hour(),
			from.Minute(),
			from.Second(),
			to.Year(),
			to.Month(),
			to.Day(),
			to.Hour(),
			to.Minute(),
			to.Second(),
		),
	)
}

func (client *Client) GetTitlesUpdatedBefore(to time.Time) ([]*Title, error) {
	return client.GetTitles(
		"/db.php",
		"Command",
		"TitleLookup",
		"TID",
		"*",
		"LastUpdate",
		fmt.Sprintf(
			"-%04d%02d%02d_%02d%02d%02d",
			to.Year(),
			to.Month(),
			to.Day(),
			to.Hour(),
			to.Minute(),
			to.Second(),
		),
	)
}

func (client *Client) GetTitlesUpdatedAfter(from time.Time) ([]*Title, error) {
	return client.GetTitles(
		"/db.php",
		"Command",
		"TitleLookup",
		"TID",
		"*",
		"LastUpdate",
		fmt.Sprintf(
			"%04d%02d%02d_%02d%02d%02d-",
			from.Year(),
			from.Month(),
			from.Day(),
			from.Hour(),
			from.Minute(),
			from.Second(),
		),
	)
}

func (client *Client) GetProgramByID(id string) (*Program, error) {
	programs, err := client.GetPrograms("Command", "ProgLookup", "id", id)
	if err != nil {
		return nil, err
	}
	if len(programs) == 0 {
		return nil, &NotFoundError{}
	}
	return programs[0], err
}

func (client *Client) GetTitles(path string, pairs ...string) ([]*Title, error) {
	data, err := client.Get(path, pairs...)
	if err != nil {
		return nil, err
	}
	return NewTitles(data)
}

func (client *Client) GetPrograms(pairs ...string) ([]*Program, error) {
	query, err := NewProgramQueryBuilder(pairs...).Build()
	if err != nil {
		return nil, err
	}
	response, err := http.Get(client.BaseURL + "/db.php" + "?" + query)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return NewPrograms(data)
}

func (client *Client) Get(path string, pairs ...string) ([]byte, error) {
	response, err := http.Get(client.BaseURL + path + "?" + createURLQueryFromKeyValue(pairs...))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

func createURLQueryFromKeyValue(pairs ...string) string {
	values := url.Values{}
	for i := 0; i < len(pairs); i += 2 {
		values.Add(pairs[i], pairs[i + 1])
	}
	return values.Encode()
}

type NotFoundError struct {}

// Returned from GetXXXByID functions.
func (error *NotFoundError) Error() string {
	return "Not Found"
}

// Utility type to convert string map to URL query.
type ProgramsQueryBuilder struct {
	Options map[string]string
}

// Creates a builder, converting key-value pairs to map[string]string.
func NewProgramQueryBuilder(pairs ...string) *ProgramsQueryBuilder {
	options := map[string]string{}
	for i := 0; i + 1 < len(pairs); i += 2 {
		options[pairs[i]] = pairs[i + 1]
	}
	return &ProgramsQueryBuilder{Options: options}
}

// Returns built query string.
func (builder *ProgramsQueryBuilder) Build() (string, error) {
	fromAndTo, err := builder.FromAndTo()
	if err != nil {
		return "", err
	}
	table := map[string]string {
		"Command": "ProgLookup",
		"Range": fromAndTo,
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

// Returns ID.
func (builder *ProgramsQueryBuilder) ID() string {
	return builder.Options["id"]
}

// Returns cal.syoboi.jp formatted :from-:to.
func (builder *ProgramsQueryBuilder) FromAndTo() (string, error) {
	from, err := builder.From()
	if err != nil {
		return "", err
	}
	to, err := builder.To()
	if err != nil {
		return "", err
	}
	if from != "" || to != "" {
		return from + "-" + to, nil
	} else {
		return "", nil
	}
}

// Returns cal.syoboi.jp formatted "from" string.
func (builder *ProgramsQueryBuilder) From() (string, error) {
	return builder.FromOrTo("from")
}

// Returns cal.syoboi.jp formatted "to" string.
func (builder *ProgramsQueryBuilder) To() (string, error) {
	return builder.FromOrTo("to")
}

// Converts RFC3339 formatted time to cal.syoboi.jp formatted time.
func (builder *ProgramsQueryBuilder) FromOrTo(key string) (string, error) {
	if str, ok := builder.Options[key]; ok {
		time, err := time.Parse(time.RFC3339, str)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(
			"%04d%02d%02d_%02d%02d%02d",
			time.Year(),
			time.Month(),
			time.Day(),
			time.Hour(),
			time.Minute(),
			time.Second(),
		), nil
	} else {
		return "", nil
	}
}
