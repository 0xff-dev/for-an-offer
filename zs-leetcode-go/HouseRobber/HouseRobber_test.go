package HouseRobber

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1}))
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{1, 2}))
}
