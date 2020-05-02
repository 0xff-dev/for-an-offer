package WordBreak

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	s1, dict1 := "leetcode", []string{"leet", "code"}
	fmt.Println(wordBreak(s1, dict1))
	s2, dict2 := "applepenapple", []string{"apple", "pen"}
	fmt.Println(wordBreak(s2, dict2))
	s3, dict3 := "catsandog", []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s3, dict3))
}