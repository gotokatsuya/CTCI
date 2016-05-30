package lesson3

import "sort"

func makeMap(text string) map[rune]bool {
	m := make(map[rune]bool, len(text))
	for _, char := range text {
		m[char] = true
	}
	return m
}

func IsMatch(rText, lText string) bool {
	if len(rText) != len(lText) {
		return false
	}
	rMap := makeMap(rText)
	lMap := makeMap(lText)
	for rChar := range rMap {
		if ok := lMap[rChar]; !ok {
			return false
		}
	}
	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func IsMatchBySort(rText, lText string) bool {
	if len(rText) != len(lText) {
		return false
	}
	rSortedText := sortString(rText)
	lSortedText := sortString(lText)
	return rSortedText == lSortedText
}
