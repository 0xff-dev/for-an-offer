package AdditiveNumber
func isAdditiveNumber(num string) bool {
    for out := 1; out < len(num); out ++ {
        for in := out+1; in < len(num); in++ {
            s1, s2 := num[0:out], num[out:in]
            // check leading zeros
            // only one zero shoule not be exclude
            if (len(s1) > 1 && s1[0] == '0') || (len(s2) > 1 && s2[0] == '0') { continue }
            next := add(s1, s2)
            cur := s1 + s2 + next
            for len(cur) < len(num) {
                s1, s2 = s2, next
                next = add(s1, s2)
                cur += next
            }
            if cur == num { return true }
        }
    }
    return false
}

func add(s1, s2 string) string {
    //  计算两个字符串相加
    maxLen := max(len(s1), len(s2))
    res := make([]uint8, 0)
    var (
        cf uint8 = 0
        s1Index uint8
        s2Index uint8
    )
    for index := 0; index < maxLen; index ++ {
        s1Index = calculate(index, &s1)
        s2Index = calculate(index, &s2)
        sum := s1Index + s2Index + cf
        cf = sum / 10
        sum %= 10
        res = append(res, sum+uint8('0'))
    }
    if cf != 0 { res = append(res, cf+uint8('0')) }
    reverse(res)
    return string(res)
}

func max(n, m int) int {
    if n > m { return n }
    return m
}

func calculate(index int, num *string) uint8 {
    if index < len(*num) {
        return (*num)[len(*num)-1-index] - uint8('0')
    }
    return 0
}

func reverse(nums []byte) {
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
