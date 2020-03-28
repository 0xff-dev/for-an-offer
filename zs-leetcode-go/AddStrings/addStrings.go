package AddStrings

// 0-48  9-57
func addStrings(num1 string, num2 string) string {
    l1, l2 := len(num1), len(num2)
    if l1 == 0 { return num2 }
    if l2 == 0 { return num1 }
    res := make([]byte, 0)
    var cf uint8 = 0
    i1, i2 := l1-1, l2-1
    for ; i1 >= 0 && i2 >= 0; i1, i2 = i1-1, i2-1 {
        a, b := num1[i1]-'0', num2[i2]-'0'
        tmp := a + b + cf
        cf = tmp / 10
        if tmp > 9 {
            tmp %= 10
        }
        res = append(res, uint8(tmp+48))
    }
    for ; i1 >= 0; i1-- {
        tmp := num1[i1]-'0' + cf
        cf = tmp / 10
        if tmp > 9 {
            tmp %= 10
        }
        res = append(res, uint8(tmp+48))
    }
    for ; i2 >= 0; i2-- {
        tmp := num2[i2] - '0' + cf
        cf = tmp / 10
        if tmp > 9 {
            tmp %= 10
        }
        res = append(res, uint8(tmp+48))
    }
    if cf != 0 { res = append(res, uint8(cf+48)) }
    for start, end := 0, len(res)-1; start < end; start, end = start+1, end-1 {
        res[start], res[end] = res[end], res[start]
    }
    return string(res)
}

