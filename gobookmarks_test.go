package gobookmarks

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	bm := Read("D:\\bak\\hmfvf\\190825.html")
	if bytes, err := bm.JSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(bytes))
	}
}
