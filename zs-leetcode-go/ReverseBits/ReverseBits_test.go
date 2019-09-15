package ReverseBits

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	for _, input := range []uint32{34520} {
		fmt.Println(reverseBits(input))
	}
}
