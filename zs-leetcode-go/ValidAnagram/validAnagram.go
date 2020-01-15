package ValidAnagram


// hash table and sort function are both good idea.
func isAnagram(s string, t string) bool {
    if len(s) != len(t) { return false }
    table := make([]int, 26)
    for index := 0; index < len(s); index++ {
        table[s[index]-'a']++
        table[t[index]-'a']--
    }
    for _, item := range table {
        if item != 0 {
            return false
        }
    }
    return true
}
