package CompareVersionNumbers

import (
    "fmt"
    "testing"
)
func TestCompareVersion(t *testing.T) {
    for _, testCase := range [][]string{
        {"0.1", "1.1"},
        {"1.0.1", "1"},
        {"7.5.2.4", "7.5.3"},
        {"1.01", "1.001"},
        {"1.0", "1.0.0"},
        {"1.2", "1.10"},
    } {
        fmt.Println(compareVersion(testCase[0], testCase[1]))
    }
}
