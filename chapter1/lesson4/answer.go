package lesson4

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

const (
	insteadOfSpace = "%20"
)

func ReplaceSpace(text string) string {
	spaceCounter := 0
	for i := len(text) - 1; i >= 0; i-- {
		r, _ := utf8.DecodeRune([]byte{text[i]})
		if space := unicode.IsSpace(r); space {
			spaceCounter++
		} else {
			break
		}
	}
	text = text[0 : len(text)-spaceCounter]
	var buffer bytes.Buffer
	for _, t := range text {
		if space := unicode.IsSpace(t); space {
			buffer.WriteString(insteadOfSpace)
		} else {
			buffer.WriteRune(t)
		}
	}
	return buffer.String()
}
