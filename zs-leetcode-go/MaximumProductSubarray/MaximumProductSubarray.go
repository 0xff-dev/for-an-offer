package MaximumProductSubarray

// 最喜爱之不需要考虑，0的问题
// 我想的N^2的算法，通过的，不考虑了
func max(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	_max := args[0]
	for _, item := range args {
		if item > _max {
			_max = item
		}
	}
	return _max
}

func next(val int) int {
	if val != 0 {
		return val
	}
	return 1
}
func maxProduct(nums []int) int {
	others := make([]int, len(nums))
	copy(others, nums)
	start, end, length := 1, len(nums)-2, len(nums)
	for ; start < length; start, end = start+1, end-1 {
		nums[start] *= next(nums[start-1])
		others[end] *= next(others[end+1])
	}
	return max(append(nums, others...)...)
}
