package ConvertSortedListToBinarySearchTree

import (
    "testing"

    "for-an-offer/zs-leetcode-go/types"
)
func TestSortedListToBST(t *testing.T) {
    list := &types.ListNode{-10, &types.ListNode{-3, &types.ListNode{0, &types.ListNode{5, &types.ListNode{9, nil}}}}}
    tree := sortedListToBST(list)
    tree.Floor()
    another := &types.ListNode{1, nil}
    anotherTree := sortedListToBST(another)
    anotherTree.Floor()
    evenList := &types.ListNode{1, &types.ListNode{2, nil}}
    evenTree := sortedListToBST(evenList)
    evenTree.Floor()
}
