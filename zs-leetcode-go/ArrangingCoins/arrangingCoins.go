package ArrangingCoins

func arrangeCoins(n int) int {
    start, cnt := 1, 0
    for n >= start {
        cnt++
        n -= start
        start ++
    }
    return cnt
}
