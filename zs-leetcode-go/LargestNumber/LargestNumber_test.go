package LargestNumber

import (
	"fmt"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	for _, item := range [][]int{{10, 2}, {3, 30, 34, 5, 9}} {
		fmt.Println(largestNumber(item))
	}
}
