package IsomorphicStrings


func isIsomorphic(s string, t string) bool {
    checkMap := make(map[byte]byte)
    store := make(map[byte]bool)// t[index] has been replace.
    for index := 0; index < len(s); index++ {
        if val, ok := checkMap[s[index]]; !ok {
            if _, ok := store[t[index]]; !ok {
                checkMap[s[index]] = t[index]
                store[t[index]] = true
            } else {
                return false
            }
        } else if val != t[index]{
            return false
        }
    }
    return true
}
