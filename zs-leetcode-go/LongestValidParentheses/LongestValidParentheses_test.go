package LongestValidParentheses

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	testStr := []string{
		"(()",
		")()())",
		"()()()",
		")()(",
		"))))",
		"((((",
		"())()()",
	}
	for _, str := range testStr {
		fmt.Println(longestValidParentheses(str))
	}
}
