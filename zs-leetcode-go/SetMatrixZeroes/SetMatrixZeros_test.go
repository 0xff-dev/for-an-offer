package SetMatrixZeroes

import (
	"fmt"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	testData := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(testData)
	fmt.Println(testData)
}
