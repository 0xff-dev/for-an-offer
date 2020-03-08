package EvaluateDivision

import (
    "fmt"
    "testing"
)

func TestCaluEquation(t *testing.T) {
    input := [][]string{
        {"a","b"}, {"b","c"},
    }
    values := []float64{2.0, 3.0}
    queries := [][]string{
        {"a","c"}, {"b","a"},
        {"a","e"}, {"a","a"},
        {"x", "x"},
    }
    fmt.Println(calcEquation(input, values, queries))
}
