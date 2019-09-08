package DecodeWays

//  取(1,2) 个 剩下的, 注意0的问题 0只可以与左边的组合,
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	for index := 1; index < len(s); index++ {
		if (s[index] == s[index-1] || s[index-1] > '2') && s[index] == '0' {
			return 0
		}
		if index == 1 {
			dp[index] = spacialForOne(s[0], s[index])
			continue
		}
		dp[index] = judge(dp, index, s[index-2], s[index-1], s[index])
	}
	return dp[len(s)-1]
}

func spacialForOne(pre, now byte) int {
	if (pre == '1' && now != '0') || (pre == '2' && now <= '6' && now != '0') {
		return 2
	}
	return 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func judge(dp []int, index int, data ...byte) int {
	if data[1] == '0' || (data[1] > '2' && data[2] != '0') || (data[1] == '2' && data[2] > '6') {
		return max(dp[index-1], dp[index-2])
	}
	if data[2] == '0' {
		return dp[index-2]
	}
	return dp[index-1] + dp[index-2]
}
