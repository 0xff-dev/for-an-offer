package FindMinimuminRotatedSortedArray

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	for _, input := range [][]int{{3, 4, 5, 1, 2}, {4, 5, 6, 7, 0, 1, 2}} {
		fmt.Println(findMin(input))
	}
}
