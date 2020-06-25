package ScoreOfParentheses

// 自己搞错了，直接用stack计算有多少对()了，这个直接利用stack.top存储当前计算的值真巧。0的时候就表明么有嵌套。
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func scoreOfParentheses(S string) int {
	stack := make([]int, 0)
	for index := 0; index < len(S); index ++ {
		length := len(stack)
		if S[index] == '(' {
			stack = append(stack, 0)
		} else {
			length -= 1
			top := stack[length]
			stack = stack[:length]
			res := max(2 * top, 1)
			if length == 0 {
				stack = []int{res}
			} else {
				stack[length-1] += res
			}
		}
	}
	return stack[0]
}
