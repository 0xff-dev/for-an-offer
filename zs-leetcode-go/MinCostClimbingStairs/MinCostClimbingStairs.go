package MinCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	if len(cost) <= 2 {
		return min(dp[0], dp[1])
	}
	for index := 2; index < len(cost); index ++ {
		dp[index] = min(dp[index-1], dp[index-2]) + cost[index]
	}

	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a, b int ) int {
	if a > b { return b}
	return a
}