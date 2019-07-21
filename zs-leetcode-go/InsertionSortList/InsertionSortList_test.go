package InsertionSortList

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	list := &types.ListNode{4, &types.ListNode{2,
		&types.ListNode{1, &types.ListNode{3, nil}}}}
	list = insertionSortList(list)
	list.Display()

	list = &types.ListNode{-1, &types.ListNode{5, &types.ListNode{3,
		&types.ListNode{4, &types.ListNode{0, nil}}}}}
	list = insertionSortList(list)
	list.Display()
}
