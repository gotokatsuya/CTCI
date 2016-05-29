package lesson2

import (
	"testing"
)

func TestReverse(t *testing.T) {
	text := "abcdefg"
	expected := "gfedcba"
	if actual := Reverse(text); actual != expected {
		t.Fatal(text)
	}
}
