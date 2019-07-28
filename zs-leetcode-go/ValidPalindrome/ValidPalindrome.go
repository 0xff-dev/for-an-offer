package ValidPalindrome

import (
	"strings"
)

func judge(char byte) bool {
	if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	start, end := 0, len(s)-1
	s = strings.ToLower(s)
	for start <= end {
		startFlag, endFlag := judge(s[start]), judge(s[end])
		if startFlag && endFlag {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		if !startFlag {
			start++
		}
		if !endFlag {
			end--
		}
	}
	return true
}