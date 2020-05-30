package UglyNumberII

import (
	"fmt"
	"testing"
)

func TestNthUglyNumbe(t *testing.T) {
	for index := 1; index <= 10; index ++ {
		fmt.Println(nthUglyNumber(index))
	}
}
