package  DeleteNodeInALinkedList

import (
    "fmt"
    "testing"
    "for-an-offer/zs-leetcode-go/types"

)

func TestDeleteNode(t *testing.T) {
    list := &types.ListNode{
        1, nil,
    }
    second := &types.ListNode{
        2, nil,
    }
    th3 := &types.ListNode {
        3, nil,
    }
    second.Next = th3
    list.Next = second
    fmt.Println(list,"  ",  second, &second, "   ", th3) // test address
    list.Display()
    tmpDeleteNode(second)
    fmt.Println(second.Val)
    list.Display()
}

func tmpDeleteNode(node *types.ListNode) {
    next := node.Next
    *node = *next
}
/*
func tmpDeleteNode(node **types.ListNode) {
    next := (*node).Next
    **node = *next
}
*/
