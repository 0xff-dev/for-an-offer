package ReverseNodesinkGroup

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	lists := &types.ListNode{1, &types.ListNode{2, &types.ListNode{3,
	&types.ListNode{4, &types.ListNode{5, &types.ListNode{6, nil}}}}}}
	root := reverseKGroup(lists, 2)
	root.Display()
	list2 := &types.ListNode{3, &types.ListNode{9, &types.ListNode{6,
	&types.ListNode{1, &types.ListNode{1, &types.ListNode{4,
	&types.ListNode{7, nil}}}}}}}
	root = reverseKGroup(list2, 4)
	root.Display()
}