package examples

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestMergeSortList(t *testing.T) {
	list1 := &types.ListNode{1, &types.ListNode{2, nil}}
	list2 := &types.ListNode{1, &types.ListNode{3, nil}}
	root := MergeSortList(list1, list2)
	root.Display()
}
