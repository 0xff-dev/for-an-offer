package SortCharactersByFrequency

import (
    "fmt"
    "testing"
)

func TestFrequencySort(t *testing.T) {
    for _, testCase := range []string{
        "tree",
        "cccaaa",
        "Aabb",
    } {
        fmt.Println(frequencySort(testCase))
    }
}
