package RemoveDuplicatesfromSortedListII

import (
	"for-an-offer/zs-leetcode-go/types"
)

func deleteDuplicates(head *types.ListNode) *types.ListNode {
	if head == nil {
		return nil
	}
	head.Display()
	mapCnt := make(map[int]int, 0)
	keys := make([]int, 0)
	for walk := head; walk != nil; walk = walk.Next {
		mapCnt[walk.Val]++
		keys = append(keys, walk.Val) // 是有顺序的
	}
	datas := make([]int, 0)
	for _, v := range keys {
		if cnt, err := mapCnt[v]; err && cnt == 1 {
			datas = append(datas, v)
		}
	}
	var pHead *types.ListNode
	var walk *types.ListNode
	for _, item := range datas {
		node := &types.ListNode{item, nil}
		if pHead == nil {
			pHead = node
			walk = node
		} else {
			walk.Next = node
		}
		walk = node
	}
	return pHead
}
