package LongestIncreasingSubsequence

import (
    "fmt"
    "testing"
)
func TestLengthOfLIS(t *testing.T) {
    for _, testCase := range [][]int{
        {10, 9, 2,5,3,7, 101, 18},
        {1,2,3,4,5,6},
    } {
        fmt.Println(lengthOfLIS(testCase))
    }
}
