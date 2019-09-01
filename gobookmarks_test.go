package gobookmarks

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	bm := Read("D:\\bak\\hmfvf\\190825.html")
	fmt.Print(bm)
}
