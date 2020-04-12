package KeyboardRow

import (
	"fmt"
	"testing"
)

func TestFindWords(t *testing.T) {
	case1 := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(case1))
}
