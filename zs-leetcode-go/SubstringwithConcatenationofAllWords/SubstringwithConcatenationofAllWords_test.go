package SubstringwithConcatenationofAllWords

import (
	"fmt"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	s := "barfoothefoobarman"
	aim := []string{"bar", "foo"}
	fmt.Println(findSubstring(s, aim))

	s = "wordgoodgoodgoodbestword"
	fmt.Println(findSubstring(s, []string{"word", "good", "best", "word"}))

	s = "wordgoodgoodgoodbestword"
	aim = []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s, aim))

	s = "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	aim = []string{"fooo", "barr", "wing", "ding", "wing"}
	fmt.Println(findSubstring(s, aim))
}
