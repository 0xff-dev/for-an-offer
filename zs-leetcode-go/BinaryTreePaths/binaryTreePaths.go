package BinaryTreePaths

import (
    "fmt"
    "for-an-offer/zs-leetcode-go/types"
)
func aux(root *types.TreeNode, now string, res *[]string) {
    if root.Left == nil && root.Right == nil { 
        *res = append(*res, now)
        return
    }
    if root.Left != nil {
        aux(root.Left, now + fmt.Sprintf("->%d", root.Left.Val), res)
    }
    if root.Right != nil {
        aux(root.Right, now + fmt.Sprintf("->%d", root.Right.Val), res)
    }
}

func binaryTreePaths(root *types.TreeNode) []string {
    if root == nil { return nil }
    res := make([]string, 0)
    now := fmt.Sprintf("%d", root.Val)
    aux(root, now, &res)
    return res
}
