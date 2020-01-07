package AddDigits
import (
    "fmt"
    "testing"
)
func TestAddDigits(t *testing.T) {
    for _, testCase := range []int{38, 238, 153, 100, 2, 10, 1234567} {
        fmt.Println(addDigits(testCase))
    }
}
