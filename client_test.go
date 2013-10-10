package sugoi

import (
	. "github.com/r7kamura/gospel"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var currentRequest *http.Request

type testHandler struct {}

func (handler *testHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	currentRequest = request
	switch request.URL.Query().Get("Command") {
	case "TitleLookup":
		fmt.Fprint(
			writer,
			`<?xml version="1.0" encoding="UTF-8"?>
			<TitleLookupResponse>
				<Result>
					<Code>200</Code>
					<Message>
					</Message>
				</Result>
				<TitleItems>
					<TitleItem id="1">
						<TID>1</TID>
						<LastUpdate>2000-01-01 00:00:00</LastUpdate>
						<Title>タイトル</Title>
						<ShortTitle></ShortTitle>
						<TitleYomi></TitleYomi>
						<TitleEN></TitleEN>
						<Comment>Comment</Comment>
						<Cat>1</Cat>
						<TitleFlag>0</TitleFlag>
						<FirstYear>2000</FirstYear>
						<FirstMonth>1</FirstMonth>
						<FirstEndYear>2000</FirstEndYear>
						<FirstEndMonth>1</FirstEndMonth>
						<FirstCh></FirstCh>
						<Keywords></Keywords>
						<UserPoint>1</UserPoint>
						<UserPointRank>1</UserPointRank>
						<SubTitles></SubTitles>
					</TitleItem>
				</TitleItems>
			</TitleLookupResponse>`,
		)
	case "ProgLookup":
		fmt.Fprint(
			writer,
			`<?xml version="1.0" encoding="UTF-8"?>
			<ProgLookupResponse>
				<ProgItems>
					<ProgItem id="1">
						<LastUpdate>2000-01-01 00:00:00</LastUpdate>
						<PID>1</PID>
						<TID>1</TID>
						<StTime>2000-01-01 00:00:00</StTime>
						<StOffset>0</StOffset>
						<EdTime>2000-01-01 00:00:00</EdTime>
						<Count>1</Count>
						<SubTitle/>
						<ProgComment/>
						<Flag>0</Flag>
						<Deleted>0</Deleted>
						<Warn>1</Warn>
						<ChID>1</ChID>
						<Revision>0</Revision>
						<STSubTitle>サブタイトル</STSubTitle>
					</ProgItem>
				</ProgItems>
				<Result>
					<Code>200</Code>
					<Message/>
				</Result>
			</ProgLookupResponse>`,
		)
	}
}

func TestClient(t *testing.T) {
	server := httptest.NewServer(&testHandler{})
	defer server.Close()
	client := NewClientWithBaseURL(server.URL)

	Describe(t, "func NewClient() *Client", func() {
		It("creates a new Client", func() {
			Expect(NewClient().BaseURL).To(Equal, "http://cal.syoboi.jp")
		})
	})

	Describe(t, "func (*Client) GetTitleByID(id string) (*Title, error)", func() {
		Context("with id", func() {
			It("sends a GET request to /db.php?Command=TitleLookup&TID=:id", func() {
				client.GetTitleByID("1")
				Expect(currentRequest.URL.Path).To(Equal, "/db.php")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=TitleLookup&TID=1")
			})
		})

		Context("with valid response", func() {
			It("returns a Title", func() {
				title, _ := client.GetTitleByID("1")
				Expect(title.CategoryID).To(Equal, "1")
				Expect(title.Comment).To(Equal, "Comment")
				Expect(title.FirstChannel).To(Equal, "")
				Expect(title.FirstEndMonth).To(Equal, "1")
				Expect(title.FirstEndYear).To(Equal, "2000")
				Expect(title.FirstMonth).To(Equal, "1")
				Expect(title.FirstYear).To(Equal, "2000")
				Expect(title.ID).To(Equal, "1")
				Expect(title.Keywords).To(Equal, "")
				Expect(title.ShortTitle).To(Equal, "")
				Expect(title.SubTitles).To(Equal, "")
				Expect(title.Title).To(Equal, "タイトル")
				Expect(title.TitleEnglish).To(Equal, "")
				Expect(title.TitleFlag).To(Equal, "0")
				Expect(title.TitleYomi).To(Equal, "")
				Expect(title.UpdatedAt).To(Equal, "2000-01-01 00:00:00")
				Expect(title.UserPoint).To(Equal, "1")
				Expect(title.UserPointRank).To(Equal, "1")
			})
		})
	})

	Describe(t, "func (*Client) GetTitles(...string) ([]*Title, error)", func() {
		Context("with updatedFrom and updatedTo", func() {
			It("sends a GET request to /db.php?Command=TitleLookup&LastUpdate=:updatedFrom-:updatedTo&TID=*", func() {
				client.GetTitles("updatedFrom", "2000-01-01T00:00:00+09:00", "updatedTo", "2000-01-02T00:00:00+09:00")
				Expect(currentRequest.URL.Path).To(Equal, "/db.php")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=TitleLookup&LastUpdate=20000101_000000-20000102_000000&TID=%2A")
			})
		})

		Context("with updatedFrom", func() {
			It("sends a GET request to /db.php?Command=TitleLookup&LastUpdate=:updatedFrom-&TID=*", func() {
				client.GetTitles("updatedFrom", "2000-01-01T00:00:00+09:00")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=TitleLookup&LastUpdate=20000101_000000-&TID=%2A")
			})
		})

		Context("with updatedTo", func() {
			It("sends a GET request to /db.php?Command=TitleLookup&LastUpdate=-:updatedTo&TID=*", func() {
				client.GetTitles("updatedTo", "2000-01-02T00:00:00+09:00")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=TitleLookup&LastUpdate=-20000102_000000&TID=%2A")
			})
		})

		Context("with id", func() {
			It("sends a GET request to /db.php?Command=TitleLookup&TID=:id", func() {
				client.GetTitles("id", "1")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=TitleLookup&TID=1")
			})
		})
	})

	Describe(t, "func (*Client) GetProgramByID(id string) (*Program, error)", func() {
		Context("with id", func() {
			It("sends a GET request to /db.php?Command=ProgLookup&PID=:id", func() {
				client.GetProgramByID("1")
				Expect(currentRequest.URL.Path).To(Equal, "/db.php")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=ProgLookup&PID=1")
			})
		})

		Context("with valid response", func() {
			It("returns a Program", func() {
				program, _ := client.GetProgramByID("1")
				Expect(program.ChannelID).To(Equal, "1")
				Expect(program.Comment).To(Equal, "")
				Expect(program.Count).To(Equal, "1")
				Expect(program.Deleted).To(Equal, "0")
				Expect(program.EditedAt).To(Equal, "2000-01-01 00:00:00")
				Expect(program.Flag).To(Equal, "0")
				Expect(program.ID).To(Equal, "1")
				Expect(program.Item).To(Equal, "")
				Expect(program.JoinedSubTitle).To(Equal, "サブタイトル")
				Expect(program.Revision).To(Equal, "0")
				Expect(program.StartedAt).To(Equal, "2000-01-01 00:00:00")
				Expect(program.StartedOffset).To(Equal, "0")
				Expect(program.SubTitle).To(Equal, "")
				Expect(program.TitleID).To(Equal, "1")
				Expect(program.UpdatedAt).To(Equal, "2000-01-01 00:00:00")
				Expect(program.Warning).To(Equal, "1")
			})
		})
	})

	Describe(t, "func (*Client) GetPrograms(...string) ([]*Program, error)", func() {
		Context("with playedFrom and playedTo", func() {
			It("sends a GET request to /db.php?Command=ProgLookup&Range=:playedFrom-:playedTo", func() {
				client.GetPrograms("playedFrom", "2000-01-01T00:00:00+09:00", "playedTo", "2000-01-02T00:00:00+09:00")
				Expect(currentRequest.URL.Path).To(Equal, "/db.php")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=ProgLookup&Range=20000101_000000-20000102_000000")
			})
		})

		Context("with playedFrom", func() {
			It("sends a GET request to /db.php?Command=ProgLookup&Range=:playedFrom-", func() {
				client.GetPrograms("playedFrom", "2000-01-01T00:00:00+09:00")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=ProgLookup&Range=20000101_000000-")
			})
		})

		Context("with playedTo", func() {
			It("sends a GET request to /db.php?Command=ProgLookup&Range=-:playedTo", func() {
				client.GetPrograms("playedTo", "2000-01-02T00:00:00+09:00")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=ProgLookup&Range=-20000102_000000")
			})
		})

		Context("with startedTo", func() {
			It("sends a GET request to /db.php?Command=ProgLookup&StTime=-:startedTo", func() {
				client.GetPrograms("startedTo", "2000-01-02T00:00:00+09:00")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=ProgLookup&StTime=-20000102_000000")
			})
		})

		Context("with channelID", func() {
			It("sends a GET request to /db.php?Command=ProgLookup&ChID=:channelID", func() {
				client.GetPrograms("channelID", "1")
				Expect(currentRequest.URL.RawQuery).To(Equal, "ChID=1&Command=ProgLookup")
			})
		})

		Context("with id", func() {
			It("sends a GET request to /db.php?Command=ProgLookup&PID=:id", func() {
				client.GetPrograms("id", "1")
				Expect(currentRequest.URL.RawQuery).To(Equal, "Command=ProgLookup&PID=1")
			})
		})
	})
}
