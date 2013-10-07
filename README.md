# Sugoi
Sugoi is a client library for [http://cal.syoboi.jp/](http://cal.syoboi.jp/) written in Golang.

## Install
```
go get github.com/r7kamura/sugoi
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/r7kamura/sugoi"
)

func main() {
	// Creates a new sugoi.Client.
	client := sugoi.NewClient()

	// Invokes func(*Client) GetTitleByID(id string) (*sugoi.Title, error).
	// The `id` is a title ID on cal.syoboi.jp, called "TID".
	title, _ := client.GetTitleByID("1")

	// sugoi.Title is a simple struct for parsed title.
	fmt.Println(
		title.Cat,           // "10"
		title.Comment,       // "..."
		title.FirstCh,       // "テレビ朝日"
		title.FirstEndMonth, // "3"
		title.FirstEndYear,  // "2003"
		title.FirstMonth,    // "1"
		title.FirstYear,     // "2003"
		title.Keywords,      // "..."
		title.LastUpdate,    // "2008-07-02 00:05:46"
		title.ShortTitle,    // ""
		title.SubTitles,     // ""
		title.TID,           // "1"
		title.Title,         // "魔法遣いに大切なこと"
		title.TitleEN,       // ""
		title.TitleFlag,     // "0"
		title.TitleYomi,     // "まほうつかいにたいせつなこと"
		title.UserPoint,     // "5"
		title.UserPointRank, // "1855"
	)
}
```
