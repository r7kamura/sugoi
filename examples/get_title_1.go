package main

import (
	".."
	"fmt"
)

func main() {
	client := sugoi.NewClient()
	title, _ := client.GetTitleByID("1")
	fmt.Println(
		title.CategoryID,
		title.Comment,
		title.FirstChannel,
		title.FirstEndMonth,
		title.FirstEndYear,
		title.FirstMonth,
		title.FirstYear,
		title.ID,
		title.Keywords,
		title.ShortTitle,
		title.SubTitles,
		title.Title,
		title.TitleEnglish,
		title.TitleFlag,
		title.TitleYomi,
		title.UpdatedAt,
		title.UserPoint,
		title.UserPointRank,
	)
}
