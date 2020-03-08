package WordLadder


func check(word string, wordList []string) bool {
    for _, str := range wordList {
        if word == str { return true }
    }
    return false
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    // first judge endword in list?
    if !check(endWord, wordList){ return 0 }
    // 数组标记？？？？？
    return 0
}
