package ExcelSheetColumnNumber

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}
