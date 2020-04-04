package FindAllAnagramsInaString

import (
    "fmt"
    "testing"
)

func TestFindAnagrams(t *testing.T) {
    case1, aim1 := "cbaebabacd", "abc"
    case2, aim2 := "abab", "ab"

    fmt.Println(findAnagrams(case1, aim1))
    fmt.Println(findAnagrams(case2, aim2))
    case3, aim3 := "baa", "aa"
    fmt.Println(findAnagrams(case3, aim3))
    case5, aim5 := `abacbabc`, `abc`
    fmt.Println(findAnagrams(case5, aim5))
}
