package WordPattern

import (
    "strings"
)
func wordPattern(pattern string, str string) bool {
    splitStr := strings.Split(str, " ")
    splitLen, strLen := len(splitStr), len(pattern)
    if splitLen != strLen { return false }
    if splitLen ==1  && strLen == 1{
        return true
    } 
    checkMap := make(map[byte]string)
    stroe := make(map[string]bool)
    for index := 0; index < strLen; index++ {
        if val, ok := checkMap[pattern[index]]; !ok {
            // only no map, we can redefine a new key.
            if _, ok := stroe[splitStr[index]]; !ok {
                checkMap[pattern[index]] = splitStr[index]
                stroe[splitStr[index]]  = true
            } else {
                return false
            }
        }else if val != splitStr[index] {
            return false
        }
    }
    return true
}
