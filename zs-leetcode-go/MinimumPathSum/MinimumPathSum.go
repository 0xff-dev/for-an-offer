package MinimumPathSum

func minPathSum(grid [][]int) int {
	lenX, lenY := len(grid), len(grid[0])
	dp := make([]int, lenY)
	for x := 0; x < lenX; x++ {
		dp[0] += grid[x][0]
		for j := 1; j < lenY; j++ {
			if dp[j] == 0 {
				dp[j] = dp[j-1] + grid[x][j]
			} else {
				if dp[j] > dp[j-1] {
					dp[j] = dp[j-1]
				}
				dp[j] += grid[x][j]
			}
		}
	}
	return dp[lenY-1]
}
