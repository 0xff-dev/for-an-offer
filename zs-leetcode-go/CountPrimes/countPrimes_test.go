package CountPrimes

import (
    "fmt"
    "testing"
)

func TestCountPrimes(t *testing.T) {
    for _, testCase := range []int{10, 5, 6, 20, 2,3} {
        fmt.Println(countPrimes(testCase))
    }
}
