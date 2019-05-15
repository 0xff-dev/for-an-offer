// 快排思想， 1做基准
package SortColors

func sortColors(nums []int) {
	length := len(nums)
	start, end := 0, length-1
	index := 0
	for ; index <= end; index++ {
		if nums[index] == 2 {
			nums[end], nums[index] = nums[index], nums[end]
			end--
			index--
		} else if nums[index] == 0 {
			nums[start], nums[index] = nums[index], nums[start]
			start++
		}
	}
}
