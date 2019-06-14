package BestTimetoBuyandSellStockII

// findMax 矩阵连城的感觉呢
// list 数字数组，now是当前的和
func maxProfit(prices []int) int {
	max := 0
	for index := 1; index < len(prices); index++ {
		if prices[index] > prices[index-1] {
			max += prices[index]-prices[index-1]
		}
	}
	return max
}
