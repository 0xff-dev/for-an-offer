package LargestRectangleinHistogram

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	// 10
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
	// 2
	fmt.Println(largestRectangleArea([]int{1,2}))
	// 4
	fmt.Println(largestRectangleArea([]int{2,2}))
	// 4
	fmt.Println(largestRectangleArea([]int{3,2,1}))
	// 12
	fmt.Println(largestRectangleArea([]int{4, 2, 6,6,2}))
	//
	fmt.Println(largestRectangleArea([]int{2,3,4,5,6}))
	fmt.Println(largestRectangleArea([]int{0}))
	fmt.Println(largestRectangleArea([]int{}))
	fmt.Println(largestRectangleArea([]int{3,3,3,5}))
	fmt.Println(largestRectangleArea([]int{2,1,2}))
}
