package PalindromePartitioning

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	for _, input := range []string{"aab", "ab", "aaa", "a", "abcd"} {
		fmt.Println(partition(input))
	}
}
