package IntersectionOfTwoArrays

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	nums1 := []int{1, 1, 2, 2}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
