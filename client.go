package sugoi

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	titles, err := client.GetTitles("id", id)
	if err != nil {
		return nil, err
	}
	if len(titles) == 0 {
		return nil, &NotFoundError{}
	}
	return titles[0], nil
}

func (client *Client) GetProgramByID(id string) (*Program, error) {
	programs, err := client.GetPrograms("id", id)
	if err != nil {
		return nil, err
	}
	if len(programs) == 0 {
		return nil, &NotFoundError{}
	}
	return programs[0], nil
}

func (client *Client) GetTitles(pairs ...string) ([]*Title, error) {
	query, err := NewTitleQueryBuilder(pairs...).Build()
	if err != nil {
		return nil, err
	}
	data, err := client.Get("/db.php", query)
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
	data, err := client.Get("/db.php", query)
	if err != nil {
		return nil, err
	}
	return NewPrograms(data)
}

func (client *Client) Get(path string, query string) ([]byte, error) {
	response, err := http.Get(client.BaseURL + path + "?" + query)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

type NotFoundError struct {}

func (error *NotFoundError) Error() string {
	return "Not Found"
}

func convertRFC3339ToSyoboiFormat(str string) (string, error) {
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
}

func convertKeyValuePairsToHash(pairs ...string) map[string]string {
	hash := map[string]string{}
	for i := 0; i + 1 < len(pairs); i += 2 {
		hash[pairs[i]] = pairs[i + 1]
	}
	return hash
}
