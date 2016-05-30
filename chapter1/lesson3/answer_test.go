package lesson3

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	rText := "abcdefg"
	lText := "gfedcba"
	if ok := IsMatch(rText, lText); !ok {
		t.Fatal(rText, lText)
	}

	rText = "abcdefg"
	lText = "gfedbba"
	if ok := IsMatch(rText, lText); ok {
		t.Fatal(rText, lText)
	}

	rText = "abcdefg"
	lText = "hgfedcba"
	if ok := IsMatch(rText, lText); ok {
		t.Fatal(rText, lText)
	}
}

func TestIsMatchBySort(t *testing.T) {
	rText := "abcdefg"
	lText := "gfedcba"
	if ok := IsMatchBySort(rText, lText); !ok {
		t.Fatal(rText, lText)
	}

	rText = "abcdefg"
	lText = "gfedbba"
	if ok := IsMatchBySort(rText, lText); ok {
		t.Fatal(rText, lText)
	}

	rText = "abcdefg"
	lText = "hgfedcba"
	if ok := IsMatchBySort(rText, lText); ok {
		t.Fatal(rText, lText)
	}
}
