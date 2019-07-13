package ReorderList

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestReorderList(t *testing.T) {
	list := &types.ListNode{1, &types.ListNode{2, &types.ListNode{3,
		&types.ListNode{4, nil}}}}
	reorderList(list)
	list.Display()
}