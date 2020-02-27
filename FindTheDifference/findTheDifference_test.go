package FindTheDifference

import (
    "fmt"
    "testing"
)

func TestFindTheDifference(t *testing.T) {
    for _, testCase := range [][]string{
        {"abcd", "abcde"},
    } {
        fmt.Println(findTheDifference(testCase[0], testCase[1]))
    }
}
