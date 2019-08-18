package MaximumProductSubarray

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	fmt.Println(maxProduct([]int{2, 3, -3, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{1}))
}
