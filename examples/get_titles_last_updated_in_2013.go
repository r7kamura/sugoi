package main

import (
	".."
	"fmt"
)

func main() {
	client := sugoi.NewClient()
	titles, _ := client.GetTitles(
		"updatedFrom", "2013-01-01T00:00:00+09:00",
		"updatedTo", "2013-02-01T00:00:00+09:00",
	)
	if len(titles) > 0 {
		fmt.Println(
			titles[0].CategoryID,
			titles[0].Comment,
			titles[0].FirstChannel,
			titles[0].FirstEndMonth,
			titles[0].FirstEndYear,
			titles[0].FirstMonth,
			titles[0].FirstYear,
			titles[0].ID,
			titles[0].Keywords,
			titles[0].ShortTitle,
			titles[0].SubTitles,
			titles[0].Title,
			titles[0].TitleEnglish,
			titles[0].TitleFlag,
			titles[0].TitleYomi,
			titles[0].UpdatedAt,
			titles[0].UserPoint,
			titles[0].UserPointRank,
		)
	} else {
		fmt.Println("Not Found")
	}
}
