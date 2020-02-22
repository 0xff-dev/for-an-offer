package IsSubsequence

func isSubsequence(s string, t string) bool {
    lens, lent := len(s), len(t)
    if lens > lent { return false }
    tIndex, sIndex := 0, 0
    for tIndex < lent && sIndex < lens {
        if s[sIndex] == t[tIndex] {
            sIndex++
        }
        tIndex++
    }
    return sIndex == lens
}
