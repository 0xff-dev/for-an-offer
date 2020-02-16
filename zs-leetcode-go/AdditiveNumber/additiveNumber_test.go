package AdditiveNumber

import (
    "fmt"
    "testing"
)

func TestIsAdditiveNumber(t *testing.T) {
    for _, testCase := range []string {
        "112358",
        "199100199",
    } {
        fmt.Println(isAdditiveNumber(testCase))
    }
}
