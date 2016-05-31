package lesson4

import (
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	text := "Mr Johan Smith    "
	expected := "Mr%20Johan%20Smith"
	if res := ReplaceSpace(text); res != expected {
		t.Fatal(text, res)
	}
}
