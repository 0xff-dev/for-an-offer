package BinaryWatch

import (
	"fmt"
	"testing"
)

func TestReadbinaryWatch(t *testing.T) {
	for _, testCase := range []int{1, 0, 2, 5} {
		fmt.Println(readBinaryWatch(testCase))
	}
}
