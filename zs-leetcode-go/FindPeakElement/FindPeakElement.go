package FindPeakElement

// 峰值元素
// 一个点成为峰值，旁边的就不会是峰值元素
func findPeakElement(nums []int) int {
	for index := 0; index < len(nums)-1; index++ {
		if nums[index] > nums[index+1] {
			return index
		}
	}
	return len(nums) - 1
}
