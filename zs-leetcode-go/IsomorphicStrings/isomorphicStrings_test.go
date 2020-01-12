package IsomorphicStrings

import (
    "fmt"
    "testing"
)
func TestIsIsomorphic(t *testing.T) {
    for _, testCase := range [][]string{
        {"egg", "add"},
        {"foo", "bar"},
        {"paper", "title"},
        {"ab", "aa"},
    }{
        fmt.Println(isIsomorphic(testCase[0], testCase[1]))
    }
}
