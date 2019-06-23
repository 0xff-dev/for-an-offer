package MaxPointsonaLine

import (
	"fmt"
	"testing"
)

func TestMaxPoints(t *testing.T) {
	data := [][]int{
		{1,1}, {2,2}, {3,3},
	}
	fmt.Println(maxPoints(data))
	data = [][]int{
		{1,1}, {3,2}, {5,3}, {4,1}, {2,3}, {1,4},
	}
	fmt.Println(maxPoints(data))
	data = [][]int{
		{0, 0}, {1,1}, {1,-1},
	}
	fmt.Println(maxPoints(data))
	data = [][]int{
		{0, 0}, {0, 0},
	}
	fmt.Println(maxPoints(data))
	data = [][]int{
		{0, 0}, {0, 0}, {1,1},
	}
	fmt.Println(maxPoints(data))
	data = [][]int{
		{1,1}, {1,1}, {1,1},
	}
	fmt.Println(maxPoints(data))
	data = [][]int{
		{1,1}, {1,1}, {2,2}, {2,2},
	}
	fmt.Println(maxPoints(data))
	fmt.Println(maxPoints([][]int{{1,1}}))
}
