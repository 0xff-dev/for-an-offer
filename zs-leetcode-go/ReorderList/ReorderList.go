package ReorderList

import "for-an-offer/zs-leetcode-go/types"


func reverse(list *types.ListNode) *types.ListNode{
	p, q := list, list.Next
	p.Next = nil
	for q != nil {
		tmp := q.Next
		q.Next = p
		p = q
		q = tmp
	}
	return p
}

func reorderList(head *types.ListNode)  {
	if head == nil {
		return
	}
	cnt := 0
	for walk := head; walk != nil; walk, cnt = walk.Next, cnt+1{}
	if cnt ==1 || cnt == 2 {
		return
	}
	var pre *types.ListNode
	walk := head
	for index := 0; index < cnt/2; index++ {
		pre = walk
		walk = walk.Next
	}
	pre.Next = nil
	walk = reverse(walk)
	res := head
	var node *types.ListNode
	for res != nil {
		node = walk
		walk = walk.Next
		node.Next = res.Next
		res.Next = node
		res = node.Next
	}
	if walk != nil {
		node.Next = walk
	}
}
