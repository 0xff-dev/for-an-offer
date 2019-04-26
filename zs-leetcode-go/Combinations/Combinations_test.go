package Combinations

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	tdata := [][]int {
		{4, 2},
		{3, 1},
	}
	for _, item := range tdata {
		fmt.Println(combine(item[0], item[1]))
	}
}