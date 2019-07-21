package InsertionSortList

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
)

func insertionSortList(head *types.ListNode) *types.ListNode {
	// 只有一个元素
	if head == nil || head.Next == nil {
		return head
	}
	// 首尾指针
	rHead := &types.ListNode{Val: head.Val, Next: nil}
	var pre, inner *types.ListNode
	for walk := head.Next; walk != nil; walk = walk.Next {
		fmt.Println(walk.Val)
		node := &types.ListNode{walk.Val, nil}
		for pre, inner = rHead, rHead; inner != nil; pre, inner = inner, inner.Next {
			if node.Val < inner.Val {
				if inner == rHead {
					node.Next = rHead
					rHead = node
				} else {
					node.Next = inner
					pre.Next = node
				}
				break
			}
		}
		if inner == nil {
			pre.Next = node
		}
	}
	return rHead
}
