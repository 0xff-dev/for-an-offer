package InterleavingString

// func aux(s1, s2, aim, res string, index1, index2 int) bool {
// 	if res == aim && index1 == len(s1) && index2 == len(s2) {
// 		return true
// 	}
// 	var flag bool
// 	if index1 < len(s1) {
// 		flag = flag || aux(s1, s2, aim, res+s1[index1:index1+1], index1+1, index2)
// 	}
// 	if index2 < len(s2) {
// 		flag = flag || aux(s1, s2, aim, res+s2[index2:index2+1], index1, index2+1)
// 	}
// 	return flag
// }
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for index := 0; index <= len(s1); index++ {
		dp[index] = make([]bool, len(s2)+1)
	}
	for out := 0; out <= len(s1); out++ {
		for innrt := 0; innrt <= len(s2); innrt++ {
			if out == 0 && innrt == 0 {
				dp[out][innrt] = true
			} else if out == 0 {
				dp[out][innrt] = dp[out][innrt-1] && s2[innrt-1] == s3[out+innrt-1]
			} else if innrt == 0 {
				dp[out][innrt] = dp[out-1][innrt] && s1[out-1] == s3[out+innrt-1]
			} else {
				dp[out][innrt] = dp[out][innrt-1] && s2[innrt-1] == s3[out+innrt-1] ||
					dp[out-1][innrt] && s1[out-1] == s3[out+innrt-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
