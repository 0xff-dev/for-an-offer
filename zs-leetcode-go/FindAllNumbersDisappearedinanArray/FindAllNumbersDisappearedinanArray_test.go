package FindAllNumbersDisappearedinanArray

import (
	"fmt"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	input := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDisappearedNumbers(input))
}
