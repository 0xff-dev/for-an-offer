package PathSumIII

import (
    "for-an-offer/zs-leetcode-go/types"
)

func pathSum(root *types.TreeNode, sum int) int {
    cnt := 0
    if root == nil { return cnt }
    queue := []*types.TreeNode{root}
    for len(queue) > 0 {
        tmp := []*types.TreeNode{}
        for _, item := range queue {
            if item.Left != nil { tmp = append(tmp, item.Left) }
            if item.Right != nil { tmp = append(tmp, item.Right) }
        }
        for _, item := range queue {
            aux(item, sum, 0, &cnt)
        }
        queue = tmp
    }
    return cnt
}

func aux(start *types.TreeNode, sum, now int, cnt *int) {
    if start == nil { return }
    if now + start.Val == sum { *cnt ++ }
    if start.Left != nil {
        aux(start.Left, sum, now + start.Val, cnt)
    }
    if start.Right != nil {
        aux(start.Right, sum, now + start.Val, cnt)
    }
}
