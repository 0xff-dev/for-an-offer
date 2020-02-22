package IsSubsequence

import (
    "fmt"
    "testing"
)


func TestIsSubsequence(t *testing.T) {
    for _, testCase := range [][]string{
        {"abc", "ahbgdc"},
        {"axc", "ahbgdc"},
        {"", "acb"},
    } {
        fmt.Println(isSubsequence(testCase[0], testCase[1]))
    }
}
