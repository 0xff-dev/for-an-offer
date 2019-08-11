package ExcelSheetColumnTitle

import (
	"fmt"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	for _, item := range []int{1, 28, 701, 52, 26, 78, 962} {
		fmt.Println(convertToTitle(item))
	}
}
