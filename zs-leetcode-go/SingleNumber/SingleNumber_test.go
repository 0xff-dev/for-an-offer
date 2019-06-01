package SingleNumber

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	fmt.Println(singleNumber([]int{2,2,1}))
	fmt.Println(singleNumber([]int{4,1,2,1,2}))
}