package NextPermutation

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testDatas := [][]int{
		[]int{1,2,3},
		[]int{3,2,1},
		[]int{1,1,5},
		[]int{1,3,2},
		[]int{1,5,6,5,3},
		[]int{5,4,7,5,3,2},
		[]int{2,2,7,5,4,3,2,2,1},
	}
	for _, data := range  testDatas {
		nextPermutation(data)
		fmt.Println(data)
	}
}
