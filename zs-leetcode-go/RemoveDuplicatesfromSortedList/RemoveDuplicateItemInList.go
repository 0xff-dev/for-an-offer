package RemoveDuplicatesfromSortedList

import "for-an-offer/zs-leetcode-go/types"

func deleteDuplicates(head *types.ListNode) *types.ListNode {
	if head == nil {
		return nil
	}
	walk, pre := head.Next, head
	pHead := head
	for walk != nil {
		if walk.Val == pre.Val {
			walk = walk.Next
			pre.Next = walk
		} else {
			walk = walk.Next
			pre = pre.Next
		}
	}
	return pHead
}