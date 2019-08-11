package LongestConsecutiveSequence

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	input := []int{100, 4, 200, 1, 2, 3}
	fmt.Println(longestConsecutive(input))
}
