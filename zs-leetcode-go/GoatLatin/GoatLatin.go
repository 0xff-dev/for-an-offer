package GoatLatin

import (
	"strings"
)

func toGoatLatin(S string) string {
	a := "a"
	ma := "ma"
	vowel := map[byte]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o' : {}, 'u': {},
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}
	words := strings.Split(S, " ")
	result := make([]string, 0)
	for _, word := range words {
		tmpWord := word
		if len(word) > 0 {
			if _, ok := vowel[word[0]]; !ok {
				tmpWord = word[1:] + word[:1]
			}
		}
		result = append(result, tmpWord+ma+a)
		a += "a"
	}
	return strings.Join(result, " ")
}
