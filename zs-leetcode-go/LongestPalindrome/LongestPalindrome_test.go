package LongestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	input := "abccccdd"
	if longestPalindrome(input) != 7 {
		t.Fatal("except 7")
	}
}
