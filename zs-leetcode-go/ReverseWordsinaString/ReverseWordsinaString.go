package ReverseWordsinaString

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	pattern := regexp.MustCompile(`\s+`)
	tmpStrArr := strings.Split(pattern.ReplaceAllString(strings.Trim(s, " "), " "), " ")
	start, end := 0, len(tmpStrArr)-1
	for ; start < end; start, end = start+1, end-1 {
		tmpStrArr[start], tmpStrArr[end] = tmpStrArr[end], tmpStrArr[start]
	}
	return strings.Join(tmpStrArr, " ")
}
