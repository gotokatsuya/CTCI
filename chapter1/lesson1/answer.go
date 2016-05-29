package lesson1

func IsUniqueChar(text string) bool {
	uniqueRuneMap := make(map[rune]bool, len(text))
	for _, char := range text {
		if ok := uniqueRuneMap[char]; ok {
			return false
		}
		uniqueRuneMap[char] = true
	}
	return true
}

func IsUniqueCharNoUseDataStructure(text string) bool {
	for li, lchar := range text {
		for ri, rChar := range text {
			if li != ri && lchar == rChar {
				return false
			}
		}
	}
	return true
}
