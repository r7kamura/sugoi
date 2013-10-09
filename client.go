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
	body, err := client.Get("/db.php", "TID", id, "Command", "TitleLookup")
	if err != nil {
		return nil, err
	}
	titles, err := NewTitles(body)
	return titles[0], err
}

func (client *Client) GetTitlesIn(from, to time.Time) ([]*Title, error) {
	data, err := client.Get(
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
	if err != nil {
		return nil, err
	}
	return NewTitles(data)
}

func (client *Client) Get(path string, pairs ...string) ([]byte, error) {
	response, err := http.Get(client.BaseURL + path + "?" + createURLQueryFromKeyValue(pairs...))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func createURLQueryFromKeyValue(pairs ...string) string {
	values := url.Values{}
	for i := 0; i < len(pairs); i += 2 {
		values.Add(pairs[i], pairs[i + 1])
	}
	return values.Encode()
}
