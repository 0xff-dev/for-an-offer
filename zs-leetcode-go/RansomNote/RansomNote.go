package RansomNote

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[byte]int)
	for _, char := range []byte(magazine) {
		magazineMap[char] ++
	}
	for _, char := range []byte(ransomNote) {
		charCount, ok := magazineMap[char]
		if !ok || charCount == 0 {
			return false
		}
		magazineMap[char]--
	}
	return true
}
