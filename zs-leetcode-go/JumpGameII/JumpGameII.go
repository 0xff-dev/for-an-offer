//Given an array of non-negative integers, you are initially positioned at the first index of the array.
//
//Each element in the array represents your maximum jump length at that position.
//
//Your goal is to reach the last index in the minimum number of jumps.

//    Input: [2,3,1,1,4]
//    Output: 2
//    Explanation: The minimum number of jumps to reach the last index is 2.
//    Jump 1 step from index 0 to 1, then 3 steps to the last index.
package JumpGameII


func max(num1, num2 int) int {
	if num1 > num2 { return num1}
	return num2
}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index, cnt := 1, 1
	now, next := nums[0], 0
	for ;index <= now && index < len(nums)-1; index++ {
		next = max(next, nums[index]+index)
		if index == now {
			cnt += 1
			now = next
		}
	}
	if index < len(nums)-1 {
		return -1
	}
	return cnt
}