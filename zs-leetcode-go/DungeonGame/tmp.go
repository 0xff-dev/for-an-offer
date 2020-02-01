package DungeonGame

// The knight's health has no upper bound.
// Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

type info struct {
    MinHP int // 所需要的最小血量
    ResHP int // 走到某个位置后，他可以剩余的血量   
}

func min(a, b info, cost int) (int, int)  {
    // a=y-1, b=y
    aDiff, bDiff := a.ResHP + cost, b.ResHP + cost
    tmpA, tmpB := info{}, info{}
    if aDiff <= 0 {
        tmpA.MinHP, tmpA.ResHP = a.MinHP - aDiff + 1, 1
    } else {
        tmpA.MinHP, tmpA.ResHP = a.MinHP,  aDiff
    }
    if bDiff <= 0 {
        tmpB.MinHP, tmpB.ResHP = b.MinHP - bDiff + 1, 1
    } else {
        tmpB.MinHP, tmpB.ResHP = b.MinHP, bDiff
    }

    if tmpA.MinHP < tmpB.MinHP { 
        return tmpA.MinHP, tmpA.ResHP
    }
    return tmpB.MinHP, tmpB.ResHP
}

func calculateMinimumHP1(dungeon [][]int) int {
    // 这是一个dp的问题
    xLen := len(dungeon)
    if xLen == 0 {
        return 0
    }
    yLen := len(dungeon[0])
    if yLen == 0 {
        return 0
    }
    dp := make([]info, yLen)
    if dungeon[0][0] < 0 {
        dp[0].MinHP, dp[0].ResHP = -dungeon[0][0]+1, 1
    } else {
        dp[0].MinHP, dp[0].ResHP = 1, dungeon[0][0]+1
    }
    for index := 1; index < yLen; index ++ {
        tmpData := dungeon[0][index]
        diff := dp[index-1].ResHP + tmpData
        if diff <= 0 {
            dp[index].MinHP, dp[index].ResHP = dp[index-1].MinHP - diff + 1, 1
        } else {
            dp[index].MinHP, dp[index].ResHP = dp[index-1].MinHP, diff 
        }
    }
    for x := 1; x < xLen; x++ {
        for y := 0; y < yLen; y++ {
            tmpData := dungeon[x][y]
            if y == 0 {
                diff := dp[y].ResHP + tmpData
                if diff <= 0 {
                    dp[y].MinHP, dp[y].ResHP = dp[y].MinHP - diff + 1, 1
                } else {
                    dp[y].ResHP = diff
                }
            } else {
                dp[y].MinHP, dp[y].ResHP = min(dp[y-1], dp[y], tmpData)
            }
        }
    }
    return dp[yLen-1].MinHP
}

