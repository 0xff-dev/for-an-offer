package PartitionList

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestPartition(t *testing.T) {
	list := &types.ListNode{4, &types.ListNode{1, &types.ListNode{3, &types.ListNode{2,
		&types.ListNode{5, &types.ListNode{2, nil}}}}}}
	partition(list, 3).Display()
	list = &types.ListNode{1, &types.ListNode{4, &types.ListNode{3, &types.ListNode{2,
		&types.ListNode{5, &types.ListNode{2, nil}}}}}}
	partition(list, 3).Display()

	list = &types.ListNode{1, nil}
	partition(list, 3).Display()
	list = &types.ListNode{1, &types.ListNode{2, &types.ListNode{3, nil}}}
	partition(list, 6).Display()
	list = &types.ListNode{6, &types.ListNode{5, &types.ListNode{1, nil}}}
	partition(list, 2).Display()
}
