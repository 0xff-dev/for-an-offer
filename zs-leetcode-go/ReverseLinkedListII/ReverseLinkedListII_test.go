package ReverseLinkedListII

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	list := &types.ListNode{Val: 1, Next: &types.ListNode{Val: 2, Next: &types.ListNode{Val: 3, Next: &types.ListNode{Val: 4, Next: &types.ListNode{Val: 5, Next: nil}}}}}
	list = reverseBetween(list, 1, 6)
	list.Display()
}
