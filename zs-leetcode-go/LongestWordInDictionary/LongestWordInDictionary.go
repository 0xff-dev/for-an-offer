package LongestWordInDictionary

import (
	"sort"
)

/*
Input:
words = ["w","wo","wor","worl", "world"]
Output: "world"
Explanation:
The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
*/

type stringSlice []string

func (ss stringSlice) Len() int {
	return len(ss)
}
func (ss stringSlice) Less(i, j int) bool {
	lenI, lenJ := len(ss[i]), len(ss[j])
	if lenI == lenJ {
		return ss[i] > ss[j]
	}
	return lenI < lenJ
}
func (ss stringSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func longestWord(words []string) string {
	sort.Sort(stringSlice(words))
	charMap := make(map[string]struct{})
	for _, word := range words {
		charMap[word] = struct{}{}
	}
	for index := len(words) - 1; index >= 0; index-- {
		found := true
		for wIndex := len(words[index]) - 1; wIndex > 0; wIndex-- {
			if _, ok := charMap[words[index][:wIndex]]; !ok {
				found = false
				break
			}
		}
		if found {
			return words[index]
		}
	}
	return ""
}
