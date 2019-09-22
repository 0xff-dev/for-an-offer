package RotateArray

func reverse(nums []int) {
	for start, end := 0, len(nums)-1; start < end; start, end = start+1, end-1 {
		nums[start], nums[end] = nums[end], nums[start]
	}
}

func rotate(nums []int, k int) {
	length := len(nums)
	if length == 0 || k == 0 {
		return
	}
	k %= length
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
