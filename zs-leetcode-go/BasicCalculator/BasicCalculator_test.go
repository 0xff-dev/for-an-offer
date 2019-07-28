package BasicCalculator

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	for _, item := range []string{"(1 + 1)", " 2-1 + 2 ", "(1+(4+5+2)-3)+(6+8)", "1-2-3", "1-1", "2147483647", "100-90+23", "1000+100-20+73"} {
		fmt.Println(calculate(item))
	}
}
