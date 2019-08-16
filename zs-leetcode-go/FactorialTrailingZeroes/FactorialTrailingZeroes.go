package FactorialTrailingZeroes

// 质因数分解？ 貌似可以，查出2，5个数, 计算0
// 得出一个结论，被2整除的概率>5
func trailingZeroes(n int) int {
	// 其实算法就是这样，算出有多少个5
	res := 0
	for n > 0 {
		res += n / 5
		n /= 5
	}
	return res
}
