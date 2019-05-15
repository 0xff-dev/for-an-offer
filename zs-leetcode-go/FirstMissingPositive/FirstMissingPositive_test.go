package FirstMissingPositive

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 5, 6, 2, 1}))
}
