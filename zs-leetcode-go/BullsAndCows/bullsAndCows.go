package BullsAndCows

import "fmt"
func getCompare(a, b int) int {
    if a <= b { return a }
    return b
}
func getHint(secret string, guess string) string {
    sMap, gMap := make(map[byte]int), make(map[byte]int)
    aCount, bCount := 0, 0
    for index := 0; index < len(secret); index++ {
        if secret[index] != guess[index] {
            sMap[secret[index]]++
            gMap[guess[index]]++
        } else {
            aCount ++
        }
    }
    for k, v := range gMap {
        if val, ok := sMap[k]; ok && val > 0{
            bCount += getCompare(v, val)
        }
    }
    return fmt.Sprintf("%dA%dB", aCount, bCount)
}
