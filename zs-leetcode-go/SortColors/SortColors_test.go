package SortColors

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	testDatas := [][]int{
		{2, 0, 2, 1, 1, 0},
		{1, 0, 0, 1, 2},
		{0},
		{2, 1},
		{2, 1, 0, 0, 2, 1},
		{1, 2, 0},
		{1, 0, 0, 0, 0, 1, 1, 2, 1},
		{2, 2, 2, 1, 0, 0, 0, 2, 1},
	}
	for _, data := range testDatas {
		sortColors(data)
		fmt.Println(data)
	}
}
