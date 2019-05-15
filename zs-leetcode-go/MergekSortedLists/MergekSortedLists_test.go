package MergekSortedLists

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestMergekLists(t *testing.T) {
	lists := []*types.ListNode{
		&types.ListNode{1, &types.ListNode{4, &types.ListNode{5, nil}}},
		&types.ListNode{1, &types.ListNode{3, &types.ListNode{4, nil}}},
		&types.ListNode{2, &types.ListNode{6, nil}},
	}
	root := mergeKLists(lists)
	root.Display()
	lists = []*types.ListNode{
		&types.ListNode{1, nil},
		&types.ListNode{0, nil},
	}
	root = mergeKLists(lists)
	root.Display()
}
