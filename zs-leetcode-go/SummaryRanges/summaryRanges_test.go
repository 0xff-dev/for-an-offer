package SummaryRanges

import (
    "fmt"
    "testing"
)

func TestSummaryRanges(t *testing.T) {
    for _, testCase := range [][]int{
        {0},
        {0, 3, 5},
        {0,1,2},
        {0,1,2,4,5,7},
        {0,2,3,4,6,8,9},
        {0,1,2,3, 5,6,7},
    } {
        fmt.Println(summaryRanges(testCase))
    }
}
