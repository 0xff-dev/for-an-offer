package IntersectionOfTwoArraysII

import (
    "fmt"
    "testing"
)

func TestIntersect(t *testing.T) {
    for _, testCase := range [][][]int {
        {{1,2,2,1}, {2,2}},
        {{4, 9, 5}, {9, 4, 9, 8, 4}},
    } {
        fmt.Println(intersect(testCase[0], testCase[1]))
    }
}
