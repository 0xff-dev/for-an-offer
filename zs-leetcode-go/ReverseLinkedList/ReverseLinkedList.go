package ReverseLinkedList

import "for-an-offer/zs-leetcode-go/types"

func reverseList(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var tail *types.ListNode
	aux(head, head.Next, &tail)
	head.Next = nil
	return tail
}

func aux(pre, next *types.ListNode, tail **types.ListNode) {
	if next.Next == nil {
		next.Next = pre
		*tail = next
		return
	}
	aux(next, next.Next, tail)
	next.Next = pre
}
