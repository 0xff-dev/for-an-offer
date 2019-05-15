package RemoveDuplicatesfromSortedArrayII

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testDatas := [][]int{
		{1, 1, 1, 2, 2, 3},
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
		{1, 2, 3},
		{1, 1, 2, 2, 3, 3},
	}
	for _, item := range testDatas {
		fmt.Println(removeDuplicates(item))
		fmt.Println(item)
	}
}
