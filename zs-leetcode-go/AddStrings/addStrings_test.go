package AddStrings

import (
    "fmt"
    "testing"
)

func TestAddStrings(t *testing.T) {
    for _, testCase := range [][]string{
        {"123", "2345"},
        {"2", "111"},
        {"999", "9901"},
        {"408", "5"},
    } {
        fmt.Println(addStrings(testCase[0], testCase[1]))
    }
}
