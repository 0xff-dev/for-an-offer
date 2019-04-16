package MinimumWindowSubstring

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s, t1 := "ADOBECODEBANC", "ABC"
	fmt.Println(minWindow(s, t1))

	s, t1 = "aaaaaaaaaaaabbbbbcdd", "abcdd"
	fmt.Println(minWindow(s, t1))
}
