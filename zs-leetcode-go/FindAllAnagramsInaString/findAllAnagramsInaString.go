package FindAllAnagramsInaString

func findAnagrams(s string, p string) []int {
    res := []int{}
    if len(s) == 0 { return res }
    pMap := make(map[byte]int8)
    pLen, sLen := len(p), len(s)
    if sLen < pLen {
        return res
    }
    for index := range p {
        pMap[p[index]]++
    }
    start, end := 0, pLen-1
    loop:
        for end < sLen {
            tmpMap := make(map[byte]int8)
            for index := start; index <= end; index ++ {
                if _, ok := pMap[s[index]]; !ok {
                    start, end = index+1, index +pLen
                    goto loop
                }
                tmpMap[s[index]]++
            }
            if compare(pMap, tmpMap) {
                res = append(res, start)
                for end < sLen-1 && s[end+1] == s[start] {
                    end, start = end+1, start+1
                    res = append(res, start)
                }
                if end < sLen -1 {
                    if _, ok := pMap[s[end+1]]; ok {
                        start, end = start+1, end+1
                        continue
                    }
                }
                start, end = end+1, end + pLen
            } else {
                start, end = start+1, end+1
            }
        }
    return res
}

func compare(map1, map2 map[byte]int8) bool {
    if len(map1) == len(map2) {
        for k, v := range map1 {
            if val, ok := map2[k]; !ok || val != v { return false }
        }
        return true
    }
    return false
}
