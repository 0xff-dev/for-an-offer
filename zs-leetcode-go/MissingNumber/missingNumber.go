package MissingNumber

func missingNumber(nums []int) int {
    length := len(nums)
    if length == 0 { return 0 }
    sum, lackSum := 0, 0
    for index := 0; index < length; index++ {
        sum += index
        lackSum += nums[index]
    }
    return sum + length - lackSum
}
