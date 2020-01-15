package ValidAnagram

import (
    "fmt"
    "testing"
)
func TestIsAnagram(t *testing.T) {
    for _, testCase := range [][]string{
        {"anagram", "nagaram"},
        {"rat","car"},
    } {
        fmt.Println(isAnagram(testCase[0], testCase[1]))
    }
}
