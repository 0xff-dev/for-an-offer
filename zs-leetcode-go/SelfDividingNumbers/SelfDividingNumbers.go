package SelfDividingNumbers

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	for start := left; start <= right; start ++ {
		mod := start
		for mod > 0 {
			_mod := mod % 10
			if _mod == 0 || start % _mod != 0 {
				break
			}
			mod /= 10
		}
		if mod == 0 {
			result = append(result, start)
		}
	}
	return result
}
