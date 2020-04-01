package ArrangingCoins

import (
    "fmt"
    "testing"
)

func TestArraningCoins(t *testing.T) {
    for _, testCase := range []int{
        1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
    } {
        fmt.Println(testCase, "=", arrangeCoins(testCase))
    }
}
