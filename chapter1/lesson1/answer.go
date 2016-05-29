package lesson1

func IsUniqueChar(text string) bool {
	m := make(map[rune]bool, len(text))
	for _, char := range text {
		if ok := m[char]; ok {
			return false
		}
		m[char] = true
	}
	return true
}

func IsUniqueCharByNotUseDataStructure(text string) bool {
	for li, lchar := range text {
		for ri, rChar := range text {
			if li != ri && lchar == rChar {
				return false
			}
		}
	}
	return true
}
