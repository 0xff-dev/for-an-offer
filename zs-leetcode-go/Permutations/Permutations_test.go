package Permutations

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	datas := [][]int{
		[]int{1, 2, 3},
		[]int{1, 4, 5},
	}
	for _, data := range datas {
		fmt.Println(permute(data))
	}
}
