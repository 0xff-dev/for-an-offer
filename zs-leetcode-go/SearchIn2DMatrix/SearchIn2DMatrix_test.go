package SearchIn2DMatrix

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	testData := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(testData, 3))
	fmt.Println(searchMatrix())
}
