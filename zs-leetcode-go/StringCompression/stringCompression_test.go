package StringCompression

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	for _, testCase := range [][]byte{
		{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		{'a'},
		{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		{'a', 'b', 'c'},
		{'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o'},
	} {
		fmt.Println(compress(testCase))
	}
}
