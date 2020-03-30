package NumberOfSegmentInaString

import (
    "fmt"
    "testing"
)

func TestCountSegments(t *testing.T) {
    for _, testCase := range []string{
        "Hello, my name is john",
        "  fdfd fdf ",
        " fdfdf fdfd",
        "",
        "fdsfsd",
    } {
        fmt.Println(countSegments(testCase))
    }
} 
