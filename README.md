# Sugoi
Sugoi client library for [http://cal.syoboi.jp/](http://cal.syoboi.jp/) written in Golang.

## Install
```
go get github.com/r7kamura/sugoi
```

## Usage
See test files for more examples.

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

## Features
Supported request examples:

* GET http://cal.syoboi.jp/db.php?Command=ProgLookup
* GET http://cal.syoboi.jp/db.php?Command=ProgLookup&JOIN=SubTitles&ChID=:channelID
* GET http://cal.syoboi.jp/db.php?Command=ProgLookup&JOIN=SubTitles&Count=:count
* GET http://cal.syoboi.jp/db.php?Command=ProgLookup&JOIN=SubTitles&PID=:id
* GET http://cal.syoboi.jp/db.php?Command=ProgLookup&JOIN=SubTitles&PID=:id
* GET http://cal.syoboi.jp/db.php?Command=ProgLookup&JOIN=SubTitles&Range=-:playedTo
* GET http://cal.syoboi.jp/db.php?Command=ProgLookup&JOIN=SubTitles&Range=:playedFrom-
* GET http://cal.syoboi.jp/db.php?Command=ProgLookup&JOIN=SubTitles&Range=:playedFrom-:playedTo
* GET http://cal.syoboi.jp/db.php?Command=ProgLookup&JOIN=SubTitles&StTime=-:startedTo
* GET http://cal.syoboi.jp/db.php?Command=TitleLookup&LastUpdate=-:updatedTo&TID=*
* GET http://cal.syoboi.jp/db.php?Command=TitleLookup&LastUpdate=:updatedFrom-&TID=*
* GET http://cal.syoboi.jp/db.php?Command=TitleLookup&LastUpdate=:updatedFrom-:updatedTo&TID=*
* GET http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=:id
