package ReverseLinkedListII

import "for-an-offer/zs-leetcode-go/types"

func reverseBetween(head *types.ListNode, m int, n int) *types.ListNode {
	if head == nil || m == n {
		return head
	}
	index, walker := 1, head
	keyList := make([]*types.ListNode, 0)
	for ; walker != nil && index <= n; index, walker = index+1, walker.Next {
		if index >= m {
			keyList = append(keyList, walker)
		}
	}
	for start, end := 0, len(keyList)-1; start < end; start, end = start+1, end-1 {
		keyList[start].Val, keyList[end].Val = keyList[end].Val, keyList[start].Val
	}
	return head
}
