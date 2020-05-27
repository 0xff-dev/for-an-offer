package LongestCommonSubsequence

import "fmt"

const keyFormat = `%d-%d`
func max(a, b int) int {
	if a > b { return a }
	return b
}
func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	if len1 == 0 || len2 == 0 {
		return 0
	}
	dp := map[string]int{}
	for t1 := 0; t1 < len1; t1 ++ {
		for t2 := 0; t2 < len2; t2 ++{
			key1, key2, key3, key4 := fmt.Sprintf(keyFormat, t1, t2),fmt.Sprintf(keyFormat, t1-1, t2-1),fmt.Sprintf(keyFormat, t1-1, t2),fmt.Sprintf(keyFormat, t1, t2-1)
			if text1[t1] == text2[t2] {
				dp[key1] = dp[key2] +1
			} else {
				dp[key1] = max(dp[key3], dp[key4])
			}
		}
	}
	return dp[fmt.Sprintf(keyFormat, len1-1, len2-1)]
}
