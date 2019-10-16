package PowerofTwo

func isPowerOfTwo(n int) bool {
	return n > 0 && (n-1)&n == 0
}

func isPowerOfTwo1(n int) bool {
	i := 1
	for i < n {
		i *= 2
	}
	return i == n
}
