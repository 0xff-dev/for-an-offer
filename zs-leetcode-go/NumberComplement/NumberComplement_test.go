package NumberComplement

import (
	"fmt"
	"testing"
)

func TestFindComplement(t *testing.T) {
	for input := 1; input < 16; input ++ {
		fmt.Println(findComplement(input))
	}
}