package ReverseString
import (
    "fmt"
    "testing"
)
func TestReverseString(t *testing.T) {
    for _, testCase := range [][]byte{
        {'h', 'e', 'l', 'l'},
    } {
        fmt.Println(testCase)
        reverseString(testCase)
        fmt.Println(testCase)
    }
}
