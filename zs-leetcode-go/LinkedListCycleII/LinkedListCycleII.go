package LinkedListCycleII

import "for-an-offer/zs-leetcode-go/types"

func detectCycle(head *types.ListNode) *types.ListNode {
	if head == nil {
		return nil
	}
	first, second := head, head.Next
	for second != nil && second.Next != nil && first != second {
		first = first.Next
		second = second.Next.Next
	}
	if first == second && first != nil {
		walk := head
		first = first.Next
		for walk != first  {
			walk = walk.Next
			first = first.Next
		}
		return walk
	}
	return nil
}