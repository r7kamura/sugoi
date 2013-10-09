package sugoi

import (
	. "github.com/r7kamura/gospel"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	var currentRequest *http.Request
	server := httptest.NewServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				currentRequest = request
				writer.Write([]byte(exampleXML))
			},
		),
	)
	defer server.Close()
	client := NewClientWithBaseURL(server.URL)

	Describe(t, "func NewClient()", func() {
		It("creates a new Client with default BaseURL", func() {
			Expect(NewClient().BaseURL).To(Equal, "http://cal.syoboi.jp")
		})
	})

	Describe(t, "func NewClientWithBaseURL(string)", func() {
		It("creates a new Client with passed BaseURL", func() {
			Expect(NewClientWithBaseURL("http://example.com").BaseURL).To(Equal, "http://example.com")
		})
	})

	Describe(t, "func (*Client) GetTitleByID(string)", func() {
		It("sends a GET request to /db.php?Command=TitleLookup&TID=:id", func() {
			client.GetTitleByID("1")
			Expect(currentRequest.URL.Path).To(Equal, "/db.php")
			Expect(currentRequest.URL.RawQuery).To(Equal, "Command=TitleLookup&TID=1")
		})
	})

	Describe(t, "func (*Client) GetTitlesIn(from, to *time.Time)", func() {
		It("sends a GET request to /db.php?Command=TitleLookup&LastUpdate=:from-:to", func() {
			jst := time.FixedZone("JST", 0)
			client.GetTitlesIn(
				time.Date(2000, 1, 1, 0, 0, 0, 0, jst),
				time.Date(2000, 1, 2, 0, 0, 0, 0, jst),
			)
			Expect(currentRequest.URL.Path).To(Equal, "/db.php")
			Expect(currentRequest.URL.RawQuery).To(Equal, "Command=TitleLookup&LastUpdate=20000101_000000-20000102_000000&TID=%2A")
		})
	})
}
