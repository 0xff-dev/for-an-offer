package BestTimetoBuyandSellStockIII

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	leftProfit := make([]int, len(prices))
	minPrice, maxPrice := prices[0], prices[len(prices)-1]
	leftMaxProfit, rightMaxProfit := 0, 0
	for index := 1; index < len(prices); index ++ {
		if prices[index] < minPrice {
			minPrice = prices[index]
		} else {
			leftMaxProfit = max(leftMaxProfit, prices[index] - minPrice)
		}
		leftProfit[index] = leftMaxProfit
	}

	res := leftProfit[len(prices)-1]
	for index := len(prices)-2; index >= 0; index -- {
		if prices[index] > maxPrice {
			maxPrice = prices[index]
		} else {
			rightMaxProfit = max(rightMaxProfit, maxPrice-prices[index])
		}
		res = max(res, rightMaxProfit + leftProfit[index])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return  a
	}
	return b
}
