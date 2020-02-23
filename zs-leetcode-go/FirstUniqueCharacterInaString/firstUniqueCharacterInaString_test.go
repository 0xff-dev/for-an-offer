package FirstUniqueCharacterInaString

import (
    "fmt"
    "testing"
)

func TestFirstUniqChar(t *testing.T) {
    for _, testCase := range []string{
        "leetcode",
        "loveleetcode",
    } {
        fmt.Println(firstUniqChar(testCase))
    }
}
