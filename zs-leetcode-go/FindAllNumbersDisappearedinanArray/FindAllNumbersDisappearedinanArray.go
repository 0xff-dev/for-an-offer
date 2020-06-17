package FindAllNumbersDisappearedinanArray

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		index := abs(num)-1
		if nums[index] > 0 {
			nums[index] *= -1
		}
	}
	for index  := range nums {
		if nums[index] > 0 {
			res = append(res, index+1)
		}
	}
	return res
}
