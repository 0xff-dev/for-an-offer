package MoveZeroes

import (
	"fmt"
	"testing"
)

func TestMoveZeroe(t *testing.T) {
	for _, data := range [][]int{{0, 1, 0, 3, 12}, {0, 0}, {1, 0, 2, 3, 0, 4}, {1}} {
		moveZeroes(data)
		fmt.Println(data)
	}
}
