package PascalTriangleII

import (
	"fmt"
	"testing"
)

func TestGetRow(t *testing.T) {
	testData := []int{0,1,2,3,4,5}
	for _, item := range testData {
		fmt.Println(getRow(item))
	}
}