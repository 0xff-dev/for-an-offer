package GrayCode

func gray(n, i int) int {
	nums := []int{0, 1, 1, 0}
	s := 0
	i *= 2
	for index := n; index > 0; index-- {
		i /= 2
		s |= nums[i%4]
		s <<= 1
	}
	return s / 2
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := make([]int, 0)
	for i := 0; i < 1<<uint(n); i++ {
		res = append(res, gray(n, i))
	}
	return res
}
