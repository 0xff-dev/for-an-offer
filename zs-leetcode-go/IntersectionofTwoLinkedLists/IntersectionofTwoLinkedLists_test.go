package IntersectionofTwoLinkedLists

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	list := &types.ListNode{4, &types.ListNode{1, &types.ListNode{8,
		&types.ListNode{4, &types.ListNode{5, nil}}}}}
	list2 := &types.ListNode{5, &types.ListNode{0, &types.ListNode{1,
		&types.ListNode{8, &types.ListNode{4, &types.ListNode{5, nil}}}}}}
	fmt.Println(getIntersectionNode(list, list2))
}