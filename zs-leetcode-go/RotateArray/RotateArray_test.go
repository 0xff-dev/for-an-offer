package RotateArray

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(array, 1)
	fmt.Println(array)
	array = []int{}
	rotate(array, 3)
	fmt.Println(array)
	array = []int{3, 2, 1}
	rotate(array, 7)
	fmt.Println(array)
	array = []int{1}
	rotate(array, 1)
	fmt.Println(array)

	array = []int{1, 2}
	rotate(array, 2)
	fmt.Println(array)
}
