package BullsAndCows

import (
    "fmt"
    "testing"
)
func TestBullsAndCows(t *testing.T) {
    for _, testCase := range [][]string{
        {"1807", "7810"},
        {"1123", "0111"},
    } {
        fmt.Println(getHint(testCase[0], testCase[1]))
    }
}
