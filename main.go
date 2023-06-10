package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/turnon/gobookmarks/bookmark"
)

func main() {
	for _, bookmarkFile := range os.Args[1:] {
		if err := bookmark.Scan(bookmarkFile, printItem); err != nil {
			panic(err)
		}
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
