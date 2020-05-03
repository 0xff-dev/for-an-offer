package WordBreakII

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	s1, dict1 := "catsanddog", []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println("ans: ", wordBreak(s1, dict1))
	for _, item := range wordBreak(s1, dict1) {
		fmt.Println("now: ", item)
	}
	s2, dict2 := "pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}
	for _, item := range wordBreak(s2, dict2) {
		fmt.Println("now: ", item)
	}
	s3, dict3 := "catsandog", []string{"cats", "dog", "sand", "and", "cat"}
	for _, item := range wordBreak(s3, dict3) {
		fmt.Println("now: ", item)
	}
	s4, dict4 := "abcd", []string{"a", "abc", "b", "cd"}
	for _, item := range wordBreak(s4, dict4) {
		fmt.Println("now4: ", item)
	}
	s5, dict5 := "aaaaaaaa", []string{"aaaa","aaa","aa"}
	fmt.Println(wordBreak(s5, dict5))
}
