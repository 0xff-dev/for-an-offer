package CountingBits


func countBits(num int) []int {
    nums := []int{0}
    for index := 1; index <= num; index++ {
        nums = append(nums, nums[index & (index-1)]+1)
    }
    return nums
}
