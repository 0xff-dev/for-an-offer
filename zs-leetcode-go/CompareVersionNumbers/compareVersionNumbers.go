package CompareVersionNumbers

import "strings"
import "strconv"
func compare(str1, str2 string) int {
    // delete leading zeros.
    num1, _ := strconv.Atoi(str1)
    num2, _ := strconv.Atoi(str2)
    if num1 > num2 {
        return 1
    } else if num1 < num2 {
        return -1
    }
    return 0
}

func isAllZero(str string) bool {
    for _, item := range str {
        if item != '0' { return false }
    }
    return true
}

func compareVersion(version1 string, version2 string) int {
    s1s := strings.Split(version1, ".")
    s2s := strings.Split(version2, ".")
    len1, len2 := len(s1s), len(s2s) 
    minLen := len1
    if len1 > len2 { minLen = len2 }
    for index := 0; index < minLen; index ++ {
        res := compare(s1s[index], s2s[index])
        if res != 0 { return res }
    }
    // 前面的都比较完成了
    // 存在的情况是如果
    if minLen == len1 {
        // 处理version2
        for ; minLen < len2; minLen ++ {
            if !isAllZero(s2s[minLen]) { return -1 }
        }
        return 0
    }
    for ; minLen < len1; minLen++ {
        if !isAllZero(s1s[minLen]) { return 1 }
    }
    return 0
}
