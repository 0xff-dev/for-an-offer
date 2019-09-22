package InterleavingString

import (
	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("a", "", "a"))
}
