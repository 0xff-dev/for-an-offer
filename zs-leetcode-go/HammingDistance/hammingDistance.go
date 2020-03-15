package HammingDistance

func hammingDistance(x int, y int) int {
    res := x ^ y
    cnt := 0
    for res > 0 {
        res = res & (res-1)
        cnt++
    }
    return cnt
}
