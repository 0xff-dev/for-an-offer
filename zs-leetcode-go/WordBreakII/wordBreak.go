package WordBreakII

func wordBreak(s string, wordDict []string) []string {
	checkMap := make(map[string][]string)
	return aux(s, wordDict, &checkMap)
}

func aux(s string, words []string, checkMap *map[string][]string) []string {
	if val, ok := (*checkMap)[s]; ok {
		return val
	}
	ans := make([]string, 0)
	if in(s, words) { ans = append(ans, s)}
	for index := 1; index < len(s); index ++ {
		word := s[index:]
		if !in(word, words) {continue}
		tmpWords := aux(s[:index], words, checkMap)
		tmpList := make([]string, 0)
		for _, item := range tmpWords {
			tmpList = append(tmpList, item+" "+word)
		}
		ans = append(ans, tmpList...)
	}
	(*checkMap)[s] = ans
	return ans
}

func in(str string, strs []string) bool {
	for _, s := range strs {
		if s == str { return true}
	}
	return false
}

//func wordBreak(s string, wordDict []string) []string {
//	res := make([]string, 0)
//	for _, word := range wordDict {
//		if strings.HasPrefix(s, word) {
//			tmp := []string{word}
//			aux(s[len(word):], wordDict, tmp, &res)
//		}
//	}
//	return res
//}
//
//func aux(s string, words []string, now []string, aim *[]string){
//	if len(s) == 0 {
//		*aim = append(*aim, strings.Join(now, " "))
//	}
//	for _, word := range words {
//		if strings.HasPrefix(s, word) {
//			aux(s[len(word):], words, append(now, word), aim)
//		}
//	}
//}
