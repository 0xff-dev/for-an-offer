package DecodeWays

import (
	"fmt"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	for _, input := range []string{"12", "226", "103", "1020", "100", "110", "130", "2130", "1050", "20"} {
		fmt.Println(numDecodings(input))
	}
}
