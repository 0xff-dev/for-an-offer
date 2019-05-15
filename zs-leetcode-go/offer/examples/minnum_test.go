package examples

import "testing"

func TestMinnum(t *testing.T) {
	nums := []int{1, 2, 3}
	Minnum(nums)
	nums = []int{23, 34, 45}
	Minnum(nums)
	nums = []int{23, 22, 1}
	Minnum(nums)
}
