package RussianDollEnvelopes

import (
    "fmt"
    "testing"
)

func TestMaxEnvelopes(t *testing.T) {
    testCase := [][]int{
        {5,4},
        {6,4},
        {6,7},
        {2,3},
    }
    fmt.Println(maxEnvelopes(testCase), " aim: ", 3)
    testCase2 := [][]int{
        {10, 8},
        {1, 12},
        {6, 15},
        {2,18},
    }
    fmt.Println(maxEnvelopes(testCase2), " aim: ", 2)
    testCase3 := [][]int{
        {4,5},
        {4,6},
        {6,7},
        {2,3},
        {1,1},
    }
    fmt.Println(maxEnvelopes(testCase3), " aim: ", 4)
    testCase4 := [][]int{
        {1,3},
        {3,5},
        {6,7},
        {6,8},
        {8,4},
        {9,5},
    }
    fmt.Println(maxEnvelopes(testCase4), " aim: ", 3)
}
