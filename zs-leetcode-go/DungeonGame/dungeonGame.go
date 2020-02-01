package DungeonGame

const inf = 0xfff
func compare(a, b int, flag bool) int {
    if a < b {
        if flag { return a }
        return b
    }
    if flag { return b }
    return a
}

// dp问题的最小化，考虑反响查找
func calculateMinimumHP(dungeon [][]int) int{
    xLen := len(dungeon)
    if xLen == 0 {
        return 0
    }
    yLen := len(dungeon[0])
    if yLen == 0 {
        return 0
    }
    dp := make([]int, yLen+1)
    // 反向递推才是正解，正向推到不仅需要考虑最小血量还需要考虑剩余血量的问题。所以那个不是正确的解法。
    for index := 0; index < yLen+1; index ++ {dp[index] = inf }
    dp[yLen-1] = 1
    for x := xLen-1; x >= 0; x-- {
        for y := yLen-1; y >= 0; y-- {
            dp[y] = compare(1, compare(dp[y], dp[y+1], true)-dungeon[x][y], false)
        }
    }
    return dp[0]
}
