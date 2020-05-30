package UglyNumberII

// usage min(min(a,b),c)
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}
	res := make([]int, n)
	res[0] = 1
	num2, num3, num5 := 0, 0, 0
	th := 1
	for th < n {
		_min := min(min(res[num2]*2, res[num3]*3), res[num5]*5)
		res[th] = _min
		for ; res[num2]*2 <= _min; num2++ {
		}
		for ; res[num3]*3 <= _min; num3++ {
		}
		for ; res[num5]*5 <= _min; num5++ {
		}
		th++
	}
	return res[n-1]
}
