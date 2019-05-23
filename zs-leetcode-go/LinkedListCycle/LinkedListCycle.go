package LinkedListCycle

import "for-an-offer/zs-leetcode-go/types"

func hasCycle(head *types.ListNode) bool {
	first, second := head, head.Next
	for second != nil && second.Next != nil && first != second {
		first = first.Next
		second = second.Next.Next
	}
	if first == second && first != nil {
		return true
	}
	return false
}