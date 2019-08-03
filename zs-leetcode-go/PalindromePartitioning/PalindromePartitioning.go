package PalindromePartitioning


func isPalindrome(s *string, start, end int) bool {
	for start < end {
		if (*s)[start] == (*s)[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}

func aux(s *string, start, end int, ans []string, res *[][]string) {
	if start == end {
		*res = append(*res, ans)
		return
	}
	for walk := start; walk < end; walk++ {
		if isPalindrome(s, start, walk){
			tmp := make([]string, len(ans))
			copy(tmp, ans)
			tmp = append(tmp, (*s)[start:walk+1])
			aux(s, walk+1, end, tmp, res)
		}
	}
}

func partition(s string) [][]string {
	answer := make([][]string, 0)
	length := len(s)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return append(answer, []string{s})
	}
	start, walk := 0, 0
	for ; walk < length; walk++ {
		if isPalindrome(&s, start, walk) {
			palindrome := make([]string, 0)
			palindrome = append(palindrome, s[start:walk+1])
			aux(&s,  walk+1, length, palindrome, &answer)
		}
	}
	return answer
}