package ThridMaximumNumber

import (
    "fmt"
    "testing"
)

func TestThridMax(t *testing.T) {
    for _, testCase := range [][]int{
        {3,2,1},
        {1,2},
        {2,2,3,1},
        {1,1,2},
    } {
        fmt.Println(thirdMax(testCase))
    }
}
