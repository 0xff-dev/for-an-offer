package RotateList

import (
	"for-an-offer/zs-leetcode-go/types"
)

// 反转连标， 前面移到后面
func reverse(head *types.ListNode) (*types.ListNode, *types.ListNode) {
	pre, walk := head, head.Next
	end := head
	pre.Next = nil
	for walk != nil {
		tmp := walk.Next
		walk.Next = pre
		pre = walk
		walk = tmp
	}
	return pre, end
}

func rotateRight(head *types.ListNode, k int) *types.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil || k == 0 {
		return head
	}
	// 尾节点
	headPoint, endPoint := reverse(head)
	for ; k > 0; k-- {
		node := headPoint
		headPoint = headPoint.Next
		endPoint.Next = node
		endPoint = node
	}
	endPoint.Next = nil
	head, _ = reverse(headPoint)
	return head
}
