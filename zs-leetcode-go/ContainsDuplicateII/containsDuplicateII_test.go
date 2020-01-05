package ContainsDuplicateII

import (
    "fmt"
    "testing"
)
func TestConatinsNearbyDuplicate(t *testing.T) {
    indexs := []int{3,1,2}
    for index, testCase := range [][]int{
        {1,2,3,1},
        {1,0,1,1},
        {1,2,3,1,2,3},
    } {
        fmt.Println(containsNearbyDuplicate(testCase, indexs[index]))
    }
}
