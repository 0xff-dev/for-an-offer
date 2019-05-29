package SearchinRotatedSortedArrayII

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	datas := []int{2,5,6,0,0,1,2}
	fmt.Println(search(datas, 0))
	fmt.Println(search(datas, 3))
	fmt.Println(search([]int{1}, 1))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{1,2}, 2))
	fmt.Println(search([]int{1,2}, 1))
	fmt.Println(search([]int{1,1}, 0))
	fmt.Println(search([]int{1,3}, 1))
	fmt.Println(search([]int{3,1}, 1))
	fmt.Println(search([]int{2,2,2,0,1}, 1))
}
