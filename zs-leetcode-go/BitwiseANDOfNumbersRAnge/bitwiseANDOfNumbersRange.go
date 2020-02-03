package BitwiseANDOfNumbersRange

// 寻找最小值的最大值，最大值的最小值
func find(num int) int {
    start := 1
    for ; start <= num; start *= 2 {}
    return start/2
}

func rangeBitwiseAnd(m int, n int) int {
    if m == n { return m }
    if m < 2 { return 0 }
    min, max := find(m), find(n)
    if min != max { return 0 }
    start := m
    for index := m+1; index <= n; index++ {
        start &= index
    }
    return start
}
