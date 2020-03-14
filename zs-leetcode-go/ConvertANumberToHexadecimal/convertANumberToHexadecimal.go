package ConvertANumberToHexadecimal


func toHex(num int) string {
    maxInt := 2147483647 // 2**31-1
    minInt := -maxInt-1 // -2**31
    compare := 4294967295 // 2**32-1
    if num > maxInt || num < minInt {
        return ""
    }
    if num == 0 {
        return "0"
    }
    if num < 0 {
        num = num+1+compare
    }
    numMap := map[int]byte{
        0:'0', 1:'1', 2:'2', 3:'3', 4:'4',
        5:'5', 6:'6', 7:'7', 8:'8', 9:'9',
        10: 'a', 11:'b', 12:'c', 13:'d', 14:'e',
        15:'f',
    }
    res := make([]byte, 0)
    for num > 0 {
        remainer := num % 16
        num = num / 16
        res = append(res, numMap[remainer])
    }
    for  s, e := 0, len(res)-1; s < e; s, e = s+1, e-1 {
        res[s], res[e] = res[e], res[s]
    }
    return string(res)
}

