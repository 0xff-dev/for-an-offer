package Triangle

import (
	"fmt"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	list := make([][]int, 0)
	list = append(list, []int{2})
	list = append(list, []int{3,4})
	list = append(list, []int{6,5,7})
	list = append(list, []int{4,1,8,3})
	fmt.Println(minimumTotal(list))
}
