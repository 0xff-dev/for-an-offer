package WordBreak

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	checkMap := make(map[int]bool)
	return aux(s, 0, wordDict, checkMap)
}

func aux(s string, index int, words []string, checkMap map[int]bool) bool {
	if index == len(s) {
		checkMap[index] = true
		return true
	}
	if val, ok := checkMap[index]; ok {
		return val
	}
	for _, word := range words {
		if strings.HasPrefix(s[index:], word) && aux(s, index+len(word), words, checkMap){
			checkMap[index] = true
			return true
		}
	}
	checkMap[index] = false
	return false
}

/*
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		var res bool
		if strings.HasPrefix(s, word) {
			res = wordBreak(s[len(word):], wordDict)
			if res { return true}
		}
	}
	return false
}
*/