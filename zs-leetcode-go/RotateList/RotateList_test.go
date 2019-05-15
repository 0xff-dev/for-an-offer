package RotateList

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestRotateRight(t *testing.T) {
	list := &types.ListNode{1, &types.ListNode{2, &types.ListNode{3,
		&types.ListNode{4, &types.ListNode{5, nil}}}}}
	root := rotateRight(list, 2)
	root.Display()

	list = &types.ListNode{0, &types.ListNode{1, &types.ListNode{2, nil}}}
	root = rotateRight(list, 4)
	root.Display()

	list = &types.ListNode{0, nil}
	root = rotateRight(list, 2)
	root.Display()

	list = &types.ListNode{}
}
