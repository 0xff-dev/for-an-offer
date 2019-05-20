package Numberof1Bits

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	for _, item := range []int{6, 3, 1} {
		fmt.Println(hammingWeight(uint32(item)))
	}
}
