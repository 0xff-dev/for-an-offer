package IntersectionofTwoLinkedLists

import (
	"for-an-offer/zs-leetcode-go/types"
)

func getIntersectionNode(headA, headB *types.ListNode) *types.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := 0, 0
	for walk := headA; walk != nil; walk, lenA = walk.Next, lenA+1{}
	for walk := headB; walk != nil; walk, lenB = walk.Next, lenB+1{}
	pa, pb := headA, headB
	if lenA > lenB {
		for index := 0; index < lenA-lenB; index++ {
			pa = pa.Next
		}
	}
	if lenA < lenB {
		for index := 0; index < lenB-lenA; index++ {
			pb = pb.Next
		}
	}
	for ;pa != nil; pa, pb = pa.Next, pb.Next {
		if pa == pb {
			return pa
		}
	}
	return nil
}
