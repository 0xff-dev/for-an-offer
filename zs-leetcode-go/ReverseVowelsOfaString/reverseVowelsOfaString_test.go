package ReverseVowelsOfaString

import (
    "fmt"
    "testing"
)
func TestReverseVowels(t *testing.T) {
    for _, testCase := range []string{
        "hello",
        "leetcode",
        "aA",
    } {
        fmt.Println(reverseVowels(testCase))
    }
}
