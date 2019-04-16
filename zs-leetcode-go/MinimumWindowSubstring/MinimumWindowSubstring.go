package MinimumWindowSubstring

// minWindow 寻找最小的窗口满足包含t的全部字母
func minWindow(s string, t string) string {
	minLength, rl, re := -1, 0, 0
	sLen, tLen := len(s), len(t)
	start, end := 0, 0
	ts := map[byte]int{}
	for i := 0; i < tLen; i++ {
		ts[t[i]]++
	}
	ans, required := 0, len(ts)
	wordCount := map[byte]int{}
	for ; end < sLen; end ++ {
		c := s[end]
		wordCount[c]++
		if cnt, err := ts[c]; err && cnt == wordCount[c] {
			ans ++
		}
		for start <= end && ans == required {
			if minLength == -1 || end-start+1 < minLength {
				minLength = end - start + 1
				rl = start
				re = end
			}
			lc := s[start]
			cnt := wordCount[lc]
			wordCount[lc] = cnt-1
			if cnt, err := ts[lc]; err && cnt > wordCount[lc] {
				ans --
			}
			start ++
		}
	}
	if minLength == -1 {
		return ""
	}
	return s[rl:re+1]
}