package lesson2

func Reverse(text string) string {
	textRunes := []rune(text)
	textRunesLen := len(textRunes)
	for i := 0; i < textRunesLen/2; i++ {
		textRunes[i], textRunes[textRunesLen-1-i] = textRunes[textRunesLen-1-i], textRunes[i]
	}
	return string(textRunes)
}
