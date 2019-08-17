package ReverseWordsinaString

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	for _, item := range []string{"the sky is blue", "  hello world!  ", "a good   example"} {
		fmt.Println(reverseWords(item))
	}
}
