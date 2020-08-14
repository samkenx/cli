package cnfjoin

import (
	"fmt"
	"testing"
)

func TestJoined(t *testing.T) {
	out, err := Joined("./testdata", "activestate", "yaml")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(out))
}
