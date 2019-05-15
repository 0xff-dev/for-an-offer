// 生成数据的全排列
package Permutations

func solve(nums []int, start int, res *[][]int) {
	if start == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		solve(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	solve(nums, 0, &res)
	return res
}
