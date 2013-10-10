package main

import (
	".."
	"fmt"
	"time"
)

func main() {
	client := sugoi.NewClient()
	jst := time.FixedZone("JST", 0)
	titles, _ := client.GetTitlesUpdatedIn(
		time.Date(2013, 1, 1, 0, 0, 0, 0, jst),
		time.Date(2013, 1, 31, 23, 59, 59, 0, jst),
	)
	if len(titles) > 0 {
		fmt.Println(
			titles[0].Cat,
			titles[0].Comment,
			titles[0].FirstCh,
			titles[0].FirstEndMonth,
			titles[0].FirstEndYear,
			titles[0].FirstMonth,
			titles[0].FirstYear,
			titles[0].Keywords,
			titles[0].LastUpdate,
			titles[0].ShortTitle,
			titles[0].SubTitles,
			titles[0].TID,
			titles[0].Title,
			titles[0].TitleEN,
			titles[0].TitleFlag,
			titles[0].TitleYomi,
			titles[0].UserPoint,
			titles[0].UserPointRank,
		)
	} else {
		fmt.Println("Not Found")
	}
}
