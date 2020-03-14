package ConvertANumberToHexadecimal

import (
    "fmt"
    "testing"
)

func TestToHex(t *testing.T) {
    for _, testCase := range []int{26, -1, 0, 1, 2147483647, -2147483648, 10, 16, 17} {
        fmt.Println(toHex(testCase))
    }

}
