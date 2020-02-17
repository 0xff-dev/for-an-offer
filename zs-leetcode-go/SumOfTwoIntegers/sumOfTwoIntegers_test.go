package SumOfTwoIntegers

import (
    "fmt"
    "testing"
)

func TestGetSum(t *testing.T) {
    for _, testCase := range [][]int {
        {78, 85},
    } {
        fmt.Println(getSum(testCase[0], testCase[1]))
    }
}
