package ReverseLinkedList

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := &types.ListNode{1, &types.ListNode{2, &types.ListNode{3, nil}}}
	reverseList(list).Display()
	reverseList(&types.ListNode{1, nil}).Display()
}
