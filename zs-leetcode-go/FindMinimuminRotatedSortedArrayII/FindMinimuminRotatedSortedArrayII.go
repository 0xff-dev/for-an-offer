package FindMinimuminRotatedSortedArrayII

func aux(nums []int, start, end int) int {
	for index := start; index < end-1; index++ {
		if nums[index] > nums[index+1] {
			return index + 1
		}
	}
	return start
}
func findMinIndex(nums []int, start, end int) int {
	minIndex := start
	for nums[start] >= nums[end] {
		if end-start == 1 {
			minIndex = end
			break
		}
		minIndex = (end + start) / 2
		if (nums[start] == nums[minIndex]) && (nums[minIndex] == nums[end]) {
			return aux(nums, start, end)
		}
		if nums[minIndex] >= nums[start] {
			start = minIndex
		} else if nums[minIndex] <= nums[end] {
			end = minIndex
		}
	}
	return minIndex
}
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[findMinIndex(nums, 0, len(nums)-1)]
}
