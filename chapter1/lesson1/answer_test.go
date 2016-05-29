package lesson1

import (
	"testing"
)

func TestIsUniqueChar(t *testing.T) {
	text := "abcdefg"
	if unique := IsUniqueChar(text); !unique {
		t.Fatal(text)
	}

	text = "abcdefgg"
	if unique := IsUniqueChar(text); unique {
		t.Fatal(text)
	}
}

func TestIsUniqueCharNoUseDataStructure(t *testing.T) {
	text := "abcdefg"
	if unique := IsUniqueCharNoUseDataStructure(text); !unique {
		t.Fatal(text)
	}

	text = "abcdefgg"
	if unique := IsUniqueCharNoUseDataStructure(text); unique {
		t.Fatal(text)
	}
}
