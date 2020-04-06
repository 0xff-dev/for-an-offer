package SortCharactersByFrequency
import (
    "sort"
    "strings"
)
type sortString []string

func (self sortString) Len() int {
    return len(self)
}
func (self sortString) Less(i, j int) bool {
    return len(self[i]) > len(self[j])
}
func (self sortString) Swap(i, j int) {
    self[i], self[j] = self[j], self[i]
}
func frequencySort(s string) string {
    charMap := make(map[byte]int)
    for index := 0; index < len(s); index ++ {
        charMap[s[index]]++
    }
    stringArray := []string{}
    for k, v := range charMap {
        stringArray = append(stringArray, generate(k, v))
    }
    sort.Sort(sortString(stringArray))
    return strings.Join(stringArray, "")
}

func generate(char byte, cnt int) string {
    bytes := []byte{}
    for count := 0; count < cnt; count++{
        bytes = append(bytes, char)
    }
    return string(bytes)
}
