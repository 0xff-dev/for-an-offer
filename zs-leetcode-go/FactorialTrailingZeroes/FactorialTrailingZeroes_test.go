package FactorialTrailingZeroes

import (
	"fmt"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	for _, input := range []int{1, 2, 3, 4, 5, 6, 10, 2147483647} {
		fmt.Println(trailingZeroes(input))
	}
}
