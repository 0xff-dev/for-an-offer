package MultiplyStrings

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	datas := [][]string{
		[]string{"2", "3"},
		[]string{"123", "456"},
		[]string{"12345", "234567"},
	}
	for _, data := range datas {
		fmt.Println(multiply(data[0], data[1]))
	}
}
