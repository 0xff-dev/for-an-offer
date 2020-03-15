package HammingDistance

import (
    "fmt"
    "testing"
)

func TestHammingDistance(t *testing.T) {
    for _, testCase := range [][]int{
        {1, 4},
        {1, 10},
    } {
        fmt.Println(hammingDistance(testCase[0], testCase[1]))
    }
}
