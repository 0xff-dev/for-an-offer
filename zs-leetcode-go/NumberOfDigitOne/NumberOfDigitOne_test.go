package NumberOfDigitOne

import (
	"fmt"
	"testing"
)

func TestCountDigitOne(t *testing.T) {
	fmt.Println(countDigitOne(13))
	fmt.Println(countDigitOne(21345))
	fmt.Println(countDigitOne(0))
	fmt.Println(countDigitOne(1))
	fmt.Println(countDigitOne(100))
}
