package LongestCommonSubsequence

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	for _, testCase := range [][]string{
		{"abcde", "ace"},
		{"abc", "abc"},
		{"abc", "def"},
	} {
		fmt.Println(longestCommonSubsequence(testCase[0], testCase[1]))
	}
}