package lesson5

import (
	"testing"
)

func TestCompress(t *testing.T) {
	text := "aabcccccddd"
	expected := "a2b1c5d3"
	if res := Compress(text); res != expected {
		t.Fatal(text, res)
	}
}
