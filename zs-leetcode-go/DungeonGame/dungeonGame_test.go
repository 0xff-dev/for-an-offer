package DungeonGame

import (
    "fmt"
    "testing"
)

func TestCalculateMinimumHP(t *testing.T) {
    testCase := [][]int{
        {-2, -3, 3},
        {-5, -10, 1},
        {10, 30, -5},
    }
    fmt.Println(calculateMinimumHP(testCase))
    testCase1 := [][]int{}
    fmt.Println(calculateMinimumHP(testCase1))
    testCase2 := [][]int{{}}
    fmt.Println(calculateMinimumHP(testCase2))
    testCase4 := [][]int{{100}}
    fmt.Println(calculateMinimumHP(testCase4))
    testCase5 := [][]int{{0, 0}}
    fmt.Println(calculateMinimumHP(testCase5))
    testCase6 := [][]int{{1, -2, 1}}
    fmt.Println(calculateMinimumHP(testCase6))
    testCase7 := [][]int{
        {1},
        {-2},
        {1},
    }
    fmt.Println(calculateMinimumHP(testCase7))
}
