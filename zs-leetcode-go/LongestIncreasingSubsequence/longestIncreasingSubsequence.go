package LongestIncreasingSubsequence

// 想到的解法如下
// 10, 9, 2, 5, 3
// 1   1  1  2  2
func lengthOfLIS(nums []int) int {
    // dp 问题
    length := len(nums)
    if length == 0 { return 0 }
    store := make([]int, length)
    maxLength := 1
    for index := 0; index < length; index++ { store[index] = 1 }
    for index := 1; index < length; index++ {
        for inner := 0; inner < index; inner++ {
            if nums[index] > nums[inner] && store[index] < store[inner]+1{
                store[index] = store[inner]+1
                if maxLength < store[index] { maxLength = store[index] }
            }
        }
    }
    return maxLength
}
