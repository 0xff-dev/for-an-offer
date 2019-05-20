package GrayCode

import (
	"fmt"
	"testing"
)

func TestGrayCode(t *testing.T) {
	testData := []int{2, 1, 0, 3}
	for _, item := range testData {
		fmt.Println(grayCode(item))
	}
}
