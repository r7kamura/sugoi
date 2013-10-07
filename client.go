package sugoi

import (
	"io/ioutil"
	"net/http"
	"net/url"
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
	title, err := NewTitle(body)
	if err != nil {
		return nil, err
	}
	return title, nil
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
