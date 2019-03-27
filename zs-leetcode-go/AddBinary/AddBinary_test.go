package AddBinary

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	fmt.Println(addBinary("1011", "11"))
	fmt.Println(addBinary("111", "11"))
	fmt.Println(addBinary("1010", "1011"))
}
