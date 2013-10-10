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
	Describe(t, "func NewTitles(string)", func() {
		It("creates new Titles by parsing XML string", func() {
			titles, _ := NewTitles([]byte(exampleXML))
			Expect(len(titles)).To(Equal, 1)
			Expect(titles[0].CategoryID).To(Equal, "1")
			Expect(titles[0].Comment).To(Equal, "Comment")
			Expect(titles[0].FirstChannel).To(Equal, "FirstCh")
			Expect(titles[0].FirstEndMonth).To(Equal, "1")
			Expect(titles[0].FirstEndYear).To(Equal, "2000")
			Expect(titles[0].FirstMonth).To(Equal, "1")
			Expect(titles[0].FirstYear).To(Equal, "2000")
			Expect(titles[0].Keywords).To(Equal, "Keywords")
			Expect(titles[0].UpdatedAt).To(Equal, "2000-01-01 00:00:00")
			Expect(titles[0].ShortTitle).To(Equal, "ShortTitle")
			Expect(titles[0].SubTitles).To(Equal, "SubTitles")
			Expect(titles[0].ID).To(Equal, "1")
			Expect(titles[0].Title).To(Equal, "タイトル")
			Expect(titles[0].TitleEnglish).To(Equal, "TitleEN")
			Expect(titles[0].TitleFlag).To(Equal, "0")
			Expect(titles[0].TitleYomi).To(Equal, "TitleYomi")
			Expect(titles[0].UserPoint).To(Equal, "1")
			Expect(titles[0].UserPointRank).To(Equal, "1")
		})
	})
}
