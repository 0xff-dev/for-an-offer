package ValidPerfectSquare

import (
    "fmt"
    "testing"
)

func TestIsPerfectSquare(t *testing.T) {
    for _, testCase := range []int{16, 14, 100, 2500, 3600, 7100, 8100} {
        fmt.Println(testCase, " ", isPerfectSquare(testCase))
    }
}
