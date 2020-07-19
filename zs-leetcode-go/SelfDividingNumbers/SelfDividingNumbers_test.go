package SelfDividingNumbers

import (
	"fmt"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	left, right := 1, 22
	fmt.Println(selfDividingNumbers(left, right))
}
