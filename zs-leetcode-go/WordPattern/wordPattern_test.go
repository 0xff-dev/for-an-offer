package WordPattern
import (
    "fmt"
    "testing"
)
func TestWordPattern(t *testing.T) {
    for _, testCase := range [][]string{
        {"abba", "dog cat cat dog"},
        {"aaa", "aa aa aa"},
        {"aaa", "a b c"},
        {"abba", "dog dog dog dog"},
    }{
        fmt.Println(wordPattern(testCase[0], testCase[1]))
    }
}
