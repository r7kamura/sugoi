package sugoi

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

var exampleXML string = `
<?xml version="1.0" encoding="UTF-8"?>
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
			<ShortTitle>ShortTitle</ShortTitle>
			<TitleYomi>TitleYomi</TitleYomi>
			<TitleEN>TitleEN</TitleEN>
			<Comment>Comment</Comment>
			<Cat>1</Cat>
			<TitleFlag>0</TitleFlag>
			<FirstYear>2000</FirstYear>
			<FirstMonth>1</FirstMonth>
			<FirstEndYear>2000</FirstEndYear>
			<FirstEndMonth>1</FirstEndMonth>
			<FirstCh>FirstCh</FirstCh>
			<Keywords>Keywords</Keywords>
			<UserPoint>1</UserPoint>
			<UserPointRank>1</UserPointRank>
			<SubTitles>SubTitles</SubTitles>
		</TitleItem>
	</TitleItems>
</TitleLookupResponse>
`

func TestTitle(t *testing.T) {
	Describe(t, "func NewTitle(string)", func() {
		It("creates a new Title by parsing XML string", func() {
			title, _ := NewTitle([]byte(exampleXML))
			Expect(title.Cat).To(Equal, "1")
			Expect(title.Comment).To(Equal, "Comment")
			Expect(title.FirstCh).To(Equal, "FirstCh")
			Expect(title.FirstEndMonth).To(Equal, "1")
			Expect(title.FirstEndYear).To(Equal, "2000")
			Expect(title.FirstMonth).To(Equal, "1")
			Expect(title.FirstYear).To(Equal, "2000")
			Expect(title.Keywords).To(Equal, "Keywords")
			Expect(title.LastUpdate).To(Equal, "2000-01-01 00:00:00")
			Expect(title.ShortTitle).To(Equal, "ShortTitle")
			Expect(title.SubTitles).To(Equal, "SubTitles")
			Expect(title.TID).To(Equal, "1")
			Expect(title.Title).To(Equal, "タイトル")
			Expect(title.TitleEN).To(Equal, "TitleEN")
			Expect(title.TitleFlag).To(Equal, "0")
			Expect(title.TitleYomi).To(Equal, "TitleYomi")
			Expect(title.UserPoint).To(Equal, "1")
			Expect(title.UserPointRank).To(Equal, "1")
		})
	})
}
