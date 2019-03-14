package MergekSortedLists

import (
	"for-an-offer/zs-leetcode-go/types"
)

// 对排序
func maintain(nodes []*types.ListNode, index, heapSize int) {
	aimIndex := index
	left, right := index*2, index*2+1
	if left <= heapSize && nodes[left].Val < nodes[aimIndex].Val {
		aimIndex = left
	}
	if right <= heapSize && nodes[right].Val < nodes[aimIndex].Val {
		aimIndex = right
	}
	if aimIndex != index {
		nodes[aimIndex], nodes[index] = nodes[index], nodes[aimIndex]
		maintain(nodes, aimIndex, heapSize)
	}
}

func buildHeap(nodes []*types.ListNode) {
	length := len(nodes)-1
	for i := len(nodes)/2; i >= 0; i-- {
		maintain(nodes, i, length)
	}
}

func mergeKLists(lists []*types.ListNode) *types.ListNode {
	nodesArr := make([]*types.ListNode, 0)
	for _, pointer := range lists {
		if pointer != nil {
			nodesArr = append(nodesArr, pointer)
		}
	}
	buildHeap(nodesArr)
	root := &types.ListNode{}
	p := root
	for len(nodesArr) > 0 {
		nextNode := nodesArr[0].Next
		nodesArr[0].Next = p.Next
		p.Next = nodesArr[0]
		p = p.Next
		if  nextNode != nil {
			nodesArr[0] = nextNode
			buildHeap(nodesArr)
		} else {
			nodesArr = nodesArr[1:]
			buildHeap(nodesArr)
		}
	}
	return root.Next
}