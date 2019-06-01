package DistinctSubsequences

import (
	"fmt"
	"testing"
)

func TestNumDistinct(t *testing.T) {
	s := "rabbbit"
	t1 := "rabbit"
	fmt.Println(numDistinct(s, t1))

	s = "babgbag"
	t1 = "bag"
	fmt.Println(numDistinct(s, t1))

	s = "rrrrr"
	t1 = "rr"
	fmt.Println(numDistinct(s, t1))

	s = "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
	t1 = "bcddceeeebecbc"
	fmt.Println(numDistinct(s, t1))
}
