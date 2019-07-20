package MaximumGap

import (
	"fmt"
	"testing"
)

func TestMaximumGap(t *testing.T) {
	data := []int{3, 6, 9, 1}
	fmt.Println(maximumGap(data))
	data = []int{10}
	fmt.Println(maximumGap(data))
}
