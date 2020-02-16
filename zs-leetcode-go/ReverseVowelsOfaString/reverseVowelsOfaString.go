package ReverseVowelsOfaString

func reverseVowels(s string) string {
    vowelsMap := map[byte]bool{
        'a': true, 'A': true,
        'e': true, 'E': true,
        'i': true, 'I': true,
        'o': true, 'O': true,
        'u': true, 'U': true,

    }
    wordBytes := []byte(s)
    i, j := 0, len(s)-1
    var (
        iFlag bool
        jFlag bool
    )
    for i < j {
        _, iFlag = vowelsMap[s[i]]
        _, jFlag = vowelsMap[s[j]]
        if iFlag && jFlag {
            wordBytes[i], wordBytes[j] = wordBytes[j], wordBytes[i]
            i, j = i+1, j-1
            continue
        }
        if !iFlag { i++  }
        if !jFlag { j-- }
    }
    return string(wordBytes)
}
