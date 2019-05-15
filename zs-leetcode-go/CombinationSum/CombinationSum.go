package CombinationSum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	stack := make([]int, 0)
	choose(candidates, target, 0, &res, stack)
	return res
}

func choose(candidates []int, target, now int, res *[][]int, stack []int) {
	for index, item := range candidates {
		now += item
		if now == target {
			tmp := make([]int, len(stack))
			copy(tmp, stack)
			*res = append(*res, append(tmp, item))
		} else if now < target {
			choose(candidates[index:], target, now, res, append(stack, item))
		}
		now -= item
	}
}
