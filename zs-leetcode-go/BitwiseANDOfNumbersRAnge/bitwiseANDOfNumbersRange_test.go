package BitwiseANDOfNumbersRange

import (
    "fmt"
    "testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
    for _, testCase := range [][]int{
        {5,7},
        {0,1},
        {1, 2},
        {0, 4},
        {1, 5},
        {3, 8},
        {6, 7},
    } {
        fmt.Println(rangeBitwiseAnd(testCase[0], testCase[1]))
    }
}
