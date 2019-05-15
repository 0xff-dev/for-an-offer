package examples

import (
	"for-an-offer/zs-leetcode-go/types"
)

func CloneList(list *types.ComplexListNode) {
	walk := list
	for walk != nil {
		cloneNode := &types.ComplexListNode{}
		cloneNode.Val = walk.Val
		cloneNode.Next = walk.Next
		cloneNode.Others = nil
		walk.Next = cloneNode
		walk = cloneNode.Next
	}
}

func ConnectOthersLink(list *types.ComplexListNode) {
	walk := list
	for walk != nil {
		cloneNode := walk.Next
		if walk.Others != nil {
			cloneNode.Others = walk.Others.Next
		}
		walk = cloneNode.Next
	}
}

func Reconnect(list *types.ComplexListNode) *types.ComplexListNode {
	pHead, walk, root := list, list.Next, list.Next
	pHead.Next = walk.Next
	pHead = pHead.Next
	for pHead != nil {
		walk.Next = pHead.Next
		walk = walk.Next
		pHead.Next = walk.Next
		pHead = pHead.Next
	}
	return root
}

func Clone(list *types.ComplexListNode) *types.ComplexListNode {
	CloneList(list)
	ConnectOthersLink(list)
	return Reconnect(list)
}
