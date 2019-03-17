package NextPermutation

import "sort"

func reverse(nums []int) {
	start, end := 0, len(nums)-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start ++
		end --
	}
}

func nextPermutation(nums []int)  {
	if len(nums) == 1 {
		return
	}
	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}
	lastIndex := len(nums)-1
	for lastIndex > 0 {
		if nums[lastIndex] <= nums[lastIndex-1]{
			lastIndex --
		} else {
			break
		}
	}
	if lastIndex == 0 {
		reverse(nums)
		return
	}
	tmp := nums[lastIndex-1]
	index := lastIndex
	for index < len(nums)-1 && nums[index+1] > tmp {
		index++
	}
	if nums[index] > tmp {
		nums[lastIndex-1], nums[index] = nums[index], nums[lastIndex-1]
	}
	sort.Ints(nums[lastIndex:])
}
