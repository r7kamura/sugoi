package main

import (
	".."
	"fmt"
)

func main() {
	client := sugoi.NewClient()
	title, _ := client.GetTitleByID("1")
	fmt.Println(
		title.Cat,
		title.Comment,
		title.FirstCh,
		title.FirstEndMonth,
		title.FirstEndYear,
		title.FirstMonth,
		title.FirstYear,
		title.Keywords,
		title.LastUpdate,
		title.ShortTitle,
		title.SubTitles,
		title.TID,
		title.Title,
		title.TitleEN,
		title.TitleFlag,
		title.TitleYomi,
		title.UserPoint,
		title.UserPointRank,
	)
}
