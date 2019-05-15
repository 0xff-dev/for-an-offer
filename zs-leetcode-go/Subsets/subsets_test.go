package Subsets

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	testDatas := [][]int{
		[]int{1, 2, 3},
		[]int{3, 2, 1},
		[]int{4, 1, 2, 5},
	}
	for _, data := range testDatas {
		fmt.Println(subsets(data))
	}
}
