package FindMinimuminRotatedSortedArrayII

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	for _, input := range [][]int{{2, 2, 2, 2, 2}, {1, 2, 3, 4, 5}, {2, 3, 4, 2, 2, 2, 2}, {2, 2, 2, 0, 1}, {3, 1, 1}} {
		fmt.Println(findMin(input))
	}
}
