package BestTimetoBuyandSellStock

func _max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	max := 0
	for out := 0; out < len(prices)-1; out++ {
		for inner := out+1; inner < len(prices); inner++ {
			if prices[inner] > prices[out] {
				max = _max(prices[inner]-prices[out], max)
			}
		}
	}
	return max
}
