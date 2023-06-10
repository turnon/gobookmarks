package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/turnon/gobookmarks/bookmark"
)

func main() {
	bookmarkFile := flag.String("file", "", "bookmark file")
	flag.Parse()

	if err := bookmark.Scan(*bookmarkFile, printItem); err != nil {
		panic(err)
	}
}

func printItem(item *bookmark.Item) error {
	itemInbytes, err := json.Marshal(item)
	if err != nil {
		return err
	}
	fmt.Println(string(itemInbytes))
	return nil
}
