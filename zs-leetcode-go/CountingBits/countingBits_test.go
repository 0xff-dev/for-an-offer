package CountingBits

import (
    "fmt"
    "testing"
)
func TestCountBits(t *testing.T) {
    for index := 0; index < 10; index ++ {
        fmt.Println(countBits(index))
    }
}
