package FractiontoRecurringDecimal

import (
	"fmt"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	fmt.Println(fractionToDecimal(9, 22))
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(-50, 8))
	fmt.Println(fractionToDecimal(-2147483648, 1))
}
