package FirstUniqueCharacterInaString

func firstUniqChar(s string) int {
    check := make(map[byte]int)
    for index := 0; index < len(s); index++ {
        check[s[index]]++
    }
    for index := 0; index < len(s); index ++ {
        if check[s[index]] == 1 {
            return index
        }
    }
    return -1
}
