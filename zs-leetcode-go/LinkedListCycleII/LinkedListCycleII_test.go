package LinkedListCycleII

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	list := &types.ListNode{3, nil}
	node := &types.ListNode{2, nil}
	node.Next = &types.ListNode{0, &types.ListNode{4, node}}
	list.Next = node
	fmt.Println(detectCycle(list).Val)
}
