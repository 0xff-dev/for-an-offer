package BestTimetoBuyandSellStockII

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{0}))
	fmt.Println(maxProfit([]int{}))
	fmt.Println(maxProfit([]int{1,0,3}))
	fmt.Println(maxProfit([]int{5,0,3}))
	fmt.Println(maxProfit([]int{3,2,6,5,0,3}))
}
