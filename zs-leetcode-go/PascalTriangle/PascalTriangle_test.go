package PascalTriangle

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println(generate(i))
	}
}
