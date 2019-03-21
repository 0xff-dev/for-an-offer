package examples

import "for-an-offer/zs-leetcode-go/types"

func MergeSortList(list1, list2 *types.ListNode) *types.ListNode{
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	root := &types.ListNode{}
	if list1.Val > list2.Val {
		root = list2
		root.Next = MergeSortList(list1, list2.Next)
	} else {
		root = list1
		root.Next = MergeSortList(list1.Next, list2)
	}
	return root
}
