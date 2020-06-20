package CheckifThereisaValidPathinaGrid

import (
	"testing"
)

func TestHasValidPath(t *testing.T) {
	// [[2,4,3],[6,5,2]] true
	input := [][]int{
		{2, 4, 3},
		{6, 5, 2},
	}
	if hasValidPath(input) != true {
		t.Fatal("[[2,4,3],[6,5,2]] except true")
	}
	// [[1,2,1],[1,2,1]] false
	input = [][]int{
		{1, 2, 1},
		{1, 2, 1},
	}
	if hasValidPath(input) != false {
		t.Fatal("[[1,2,1],[1,2,1]] except false")
	}
	// [[1,1,2]] false
	input = [][]int{
		{1, 1, 2},
	}
	if hasValidPath(input) != false {
		t.Fatal("[[1,1,2]]  except false")
	}
	// [[1,1,1,1,1,1,3]] true
	input = [][]int{
		{1, 1, 1, 1, 1, 1, 3},
	}
	if hasValidPath(input) != true {
		t.Fatal("[[1,1,1,1,1,1,3]] except true")
	}
	// [[2],[2],[2],[2],[2],[2],[6]] true
	input = [][]int{
		{2},
		{2},
		{2},
		{2},
		{2},
		{2},
		{6},
	}
	if hasValidPath(input) != true {
		t.Fatal("[[2],[2],[2],[2],[2],[2],[6]] except true")
	}
}
