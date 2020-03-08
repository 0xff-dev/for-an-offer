package SumOfLeftLeaves

import "for-an-offer/zs-leetcode-go/types"

func sumOfLeftLeaves(root *types.TreeNode) int {
    result := 0
    aux(root, false, &result)
    return result
}

func aux(root *types.TreeNode, flag bool, now *int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil && flag {
        *now = *now + root.Val
        return
    }
    if root.Left != nil {
        aux(root.Left, true, now)
    }
    if root.Right != nil {
        aux(root.Right, false, now)
    }
}
