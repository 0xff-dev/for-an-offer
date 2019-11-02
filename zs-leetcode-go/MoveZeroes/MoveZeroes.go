package MoveZeroes

// [0 1 0 3 12] 维持第一个0的位置，一个一个交换
// 1 0 0 3 12
// 0 1 2 3 0 0 3
// 1 0 2 3  0 0 3
//  1 2 3 0
func moveZeroes(nums []int) {
	if nums == nil {
		return
	}
	firstZeroIndex := 0
	for ; firstZeroIndex < len(nums) && nums[firstZeroIndex] != 0; firstZeroIndex++ {
	}
	for index := firstZeroIndex + 1; index < len(nums); index++ {
		if nums[index] != 0 {
			nums[firstZeroIndex] = nums[index]
			nums[index] = 0
			firstZeroIndex++
		}
	}
}
