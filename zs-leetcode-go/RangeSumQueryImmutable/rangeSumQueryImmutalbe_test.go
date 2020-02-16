package RangeSumQueryImmutable

import (
    "fmt"
    "testing"
)

func TestRange(t *testing.T) {
    obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
    for _, testCase := range [][]int{
        {0, 2},
        {2, 5},
        {0, 5},
    } {
        fmt.Println(obj.SumRange(testCase[0], testCase[1]))
    }
}
