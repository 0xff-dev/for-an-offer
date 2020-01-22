package ConvertSortedListToBinarySearchTree

import "for-an-offer/zs-leetcode-go/types"

func aux(root *types.TreeNode, flag bool, list *types.ListNode, end *types.ListNode) {
    if list == end {
        return
    }
    pre, walker := list, list
    for walker != end && walker.Next != end {
        pre = pre.Next
        walker = walker.Next.Next
    }
    // pre is the mid node in the list
    treeNode := &types.TreeNode{
        Val: pre.Val, 
        Left: nil, 
        Right: nil,
    }
    aux(treeNode, false, pre.Next, end)
    aux(treeNode, true, list, pre)
    if flag {
        root.Left = treeNode
    } else {
        root.Right = treeNode
    }
}
func sortedListToBST(head *types.ListNode) *types.TreeNode {
    if head == nil {
        return nil
    }
    pre, walker := head, head
    for walker != nil && walker.Next != nil {
        pre = pre.Next
        walker = walker.Next.Next
    }
    treeNode := &types.TreeNode{
        Val: pre.Val,
        Left: nil,
        Right: nil,
    }
    aux(treeNode, false, pre.Next, nil)
    aux(treeNode, true, head, pre)
    return treeNode
}
