package SingleNumberII

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	for _, arr := range [][]int{{1, 1, 1, 3}, {0, 1, 0, 1, 0, 1, 99}, {2, 2, 3, 2}} {
		fmt.Println(singleNumber(arr))
	}
}
