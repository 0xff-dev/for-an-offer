package ClimbingStairs

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	testDatas := []int{1,2,3,4,5,6,7,8,9,10}
	for _, item := range testDatas{
		fmt.Println(climbStairs(item))
	}
}
