package MinCostClimbingStairs

import (
	"fmt"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
	cost = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
