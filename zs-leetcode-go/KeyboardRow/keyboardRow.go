package KeyboardRow

import "strings"

// 97-122
const (
	one   = `qwertyuiop`
	two   = `asdfghjkl`
	three = `zxcvbnm`
)

func findWords(words []string) []string {
	result := make([]string, 0)
	checkMap := [26]byte{}
	for _, char := range one {
		checkMap[char-'a'] = 1
	}
	for _, char := range two {
		checkMap[char-'a'] = 2
	}
	for _, char := range three {
		checkMap[char-'a'] = 3
	}
	for _, word := range words {
		if word == "" {
			result = append(result, word)
			continue
		}
		tmpword := strings.ToLower(word)
		first, index := checkMap[tmpword[0]-'a'], 1
		for ; index < len(word); index++ {
			if checkMap[tmpword[index]-'a'] != first {
				break
			}
		}
		if index == len(word) {
			result = append(result, word)
		}
	}
	return result
}
