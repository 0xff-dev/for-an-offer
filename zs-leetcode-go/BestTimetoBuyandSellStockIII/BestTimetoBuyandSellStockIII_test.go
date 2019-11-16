package BestTimetoBuyandSellStockIII

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	for _, test := range [][]int{{3, 3, 5, 0, 0, 3, 1, 4}} {
		fmt.Println(maxProfit(test))
	}
}