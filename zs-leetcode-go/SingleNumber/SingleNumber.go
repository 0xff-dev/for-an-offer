package SingleNumber


func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for index := 1; index < len(nums); index++ {
		res ^= nums[index]
	}
	return res
}
