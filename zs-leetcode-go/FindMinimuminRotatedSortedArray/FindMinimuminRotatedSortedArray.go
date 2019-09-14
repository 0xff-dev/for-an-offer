package FindMinimuminRotatedSortedArray

// O(n) 遍历找到转折点
// 二分可以考虑
func findMin(nums []int) int {
	start, end, minIndex := 0, len(nums)-1, 0
	for nums[start] > nums[end] {
		if end-start == 1 {
			minIndex = end
			break
		}
		minIndex = (end + start) / 2
		if nums[minIndex] < nums[end] {
			end = minIndex
		} else if nums[minIndex] > nums[start] {
			start = minIndex
		}
	}
	return nums[minIndex]
}
