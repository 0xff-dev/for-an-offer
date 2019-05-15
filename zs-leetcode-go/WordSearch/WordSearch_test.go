package WordSearch

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
	fmt.Println(exist([][]byte{}, "dfg"))

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCESEEEFS"))
}
