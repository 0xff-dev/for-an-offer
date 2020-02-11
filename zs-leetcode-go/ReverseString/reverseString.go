package ReverseString

func reverseString(s []byte) {
    if len(s) == 0 { return }
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left, right = left+1, right-1
    }
}
