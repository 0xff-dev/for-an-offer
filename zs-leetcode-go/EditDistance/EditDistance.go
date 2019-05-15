package EditDistance

// minDistance 寻找两个字符串的最小更更改次数完成匹配。
// 也是自己第一次去接触这道题， 之前只是看看。
func minDistance(word1 string, word2 string) int {
	w1, w2 := len(word1), len(word2)
	if len(word2) == 0 {
		return w1
	}
	if len(word1) == 0 {
		return w2
	}
	dp := make([][]int, w1+1)
	for i := 0; i < w1+1; i++ {
		dp[i] = make([]int, w2+1)
	}
	for i := 0; i < w1+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < w2+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= w1; i++ {
		for j := 1; j <= w2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[w1][w2]
}

func min(args ...int) int {
	minNum := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < minNum {
			minNum = args[i]
		}
	}
	return minNum
}
