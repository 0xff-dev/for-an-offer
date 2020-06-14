package FizzBuzz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for testCase := 0; testCase < 16; testCase ++ {
		fmt.Println(fizzBuzz(testCase))
	}
}