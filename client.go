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

func (client *Client) GetTitlesIn(from, to time.Time) ([]*Title, error) {
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

func (client *Client) GetTitlesBefore(to time.Time) ([]*Title, error) {
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

func (client *Client) GetTitlesAfter(from time.Time) ([]*Title, error) {
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
	programs, err := client.GetPrograms("/db.php", "Command", "ProgLookup", "PID", id)
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

func (client *Client) GetPrograms(path string, pairs ...string) ([]*Program, error) {
	data, err := client.Get(path, pairs...)
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

func (error *NotFoundError) Error() string {
	return "Not Found"
}
