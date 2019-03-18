package sqrtx

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	datas := []int{
		4, 8, 10, 36,
	}
	for _, data := range  datas {
		fmt.Println(mySqrt(data))
	}
}
