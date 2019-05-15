package CombinationSum

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(candidates, 7))
	candidates = []int{2, 3, 5}
	fmt.Println(combinationSum(candidates, 8))
}
