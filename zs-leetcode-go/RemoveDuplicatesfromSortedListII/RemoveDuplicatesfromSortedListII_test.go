package RemoveDuplicatesfromSortedListII

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	list := &types.ListNode{1, &types.ListNode{2, &types.ListNode{3, &types.ListNode{3,
		&types.ListNode{4, &types.ListNode{4, &types.ListNode{5, nil}}}}}}}
	deleteDuplicates(list).Display()
	list = &types.ListNode{1, &types.ListNode{1, &types.ListNode{1,
		&types.ListNode{2, &types.ListNode{3, nil}}}}}
	deleteDuplicates(list).Display()
	list = &types.ListNode{2, &types.ListNode{1, nil}}
	deleteDuplicates(list).Display()
	list = &types.ListNode{-2147483648, &types.ListNode{2147483647,&types.ListNode{2, nil}}}
	deleteDuplicates(list).Display()

}
