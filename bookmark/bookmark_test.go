package bookmark

import (
	"testing"
)

func TestRead(t *testing.T) {
	Scan("D:\\bak\\hmfvf\\190825.html", func(item *Item) error {
		t.Log(t)
		return nil
	})
}
