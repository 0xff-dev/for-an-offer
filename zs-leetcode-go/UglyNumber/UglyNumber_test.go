package UglyNumber

import (
    "fmt"
    "testing"
)
func TestIsUgly(t *testing.T) {
    for _, data := range []int{6,8,14} {
        fmt.Println(isUgly(data))
    }
}
