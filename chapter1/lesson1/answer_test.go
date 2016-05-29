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

func TestIsUniqueCharByNotUseDataStructure(t *testing.T) {
	text := "abcdefg"
	if unique := IsUniqueCharByNotUseDataStructure(text); !unique {
		t.Fatal(text)
	}
	text = "abcdefgg"
	if unique := IsUniqueCharByNotUseDataStructure(text); unique {
		t.Fatal(text)
	}
}
