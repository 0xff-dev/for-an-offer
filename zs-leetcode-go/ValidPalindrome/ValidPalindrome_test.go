package ValidPalindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	data := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(data))
	data = "race a car"
	fmt.Println(isPalindrome(data))
}
