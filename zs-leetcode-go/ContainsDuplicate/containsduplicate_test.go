package ContainsDuplicate

import (
    "fmt"
    "testing"
)
func TestContaineDuplicate(t *testing.T) {
    for _, testCase := range [][]int{
        {1,2,3,1},
        {1,2,3,4},
        {1,1,1,3,3,4,3,2,4,2},
    } {
        fmt.Println(containsDuplicate(testCase))
    }
}
