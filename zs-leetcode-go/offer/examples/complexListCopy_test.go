package examples

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestClone(t *testing.T) {
	list4 := &types.ComplexListNode{4, nil, nil}
	list3 := &types.ComplexListNode{3, list4, nil}
	list2 := &types.ComplexListNode{2, list3, list4}
	list1 := &types.ComplexListNode{1, list2, list4}
	root := Clone(list1)
	root.Display()
}
