package RemoveLinkedListelements

import "testing"

import "for-an-offer/zs-leetcode-go/types"

func TestRemoveElements(t *testing.T) {
	list := &types.ListNode{1, &types.ListNode{2, &types.ListNode{3, &types.ListNode{4, &types.ListNode{5, &types.ListNode{6, nil}}}}}}
	removeElements(list, 6)
	list.Display()
}
