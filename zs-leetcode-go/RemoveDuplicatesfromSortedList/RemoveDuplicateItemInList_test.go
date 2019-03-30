package RemoveDuplicatesfromSortedList

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func auxTest(list *types.ListNode) {
	root := deleteDuplicates(list)
	root.Display()
}
func TestDeleteDuplicates(t *testing.T) {
	list := &types.ListNode{1, &types.ListNode{1, &types.ListNode{2, nil}}}
	auxTest(list)
	list = &types.ListNode{1, &types.ListNode{1, &types.ListNode{2,
		&types.ListNode{3, &types.ListNode{3, nil}}}}}
	auxTest(list)
	list = &types.ListNode{1, &types.ListNode{1, &types.ListNode{1, nil}}}
	auxTest(list)
}
