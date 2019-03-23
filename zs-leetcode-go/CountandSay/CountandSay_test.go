package CountandSay

import (
	"fmt"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	testDatas := []int{5,6,7,8,9}
	for _, data := range testDatas {
		fmt.Println(countAndSay(data))
	}
}
