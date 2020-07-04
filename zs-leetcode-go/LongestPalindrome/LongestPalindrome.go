package LongestPalindrome

func longestPalindrome(s string) int {
	charCount := map[byte]int{}
	for _, char := range s {
		charCount[byte(char)]++
	}
	length := 0
	addOne := false
	for _, v := range charCount {
		if v % 2 == 0 {
			length += v
		} else {
			length += v-1
			addOne = true
		}
	}
	if addOne {
		length+=1
	}
	return length
}
