package sugoi

import (
	. "github.com/r7kamura/gospel"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
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
		var currentRequest *http.Request
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			currentRequest = request
			writer.Write([]byte(exampleXML))
		}))
		defer server.Close()
		client := NewClient()
		client.BaseURL = server.URL

		It("sends a GET request to /db.php?Command=TitleLookup&TID=:id", func() {
			client.GetTitleByID("1")
			Expect(currentRequest.URL.Path).To(Equal, "/db.php")
			Expect(currentRequest.URL.RawQuery).To(Equal, "Command=TitleLookup&TID=1")
		})

		It("returns a Title by parsing XML response", func() {
			title, _ := client.GetTitleByID("1")
			Expect(title.Title).To(Equal, "タイトル")
		})
	})
}
