package RemoveLinkedListelements

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
)

func removeElements(head *types.ListNode, val int) *types.ListNode {
	// 直接利用双重指针完成即可
	if head == nil {
		return nil
	}
	walker := &head
	for walker != nil {
		entry := *walker
		fmt.Println(entry.Val)
		if entry.Val == val {
			*walker = entry.Next
			if entry.Next == nil {
				walker = nil
			}
		} else {
			if entry.Next == nil {
				walker = nil
			} else {
				walker = &(entry.Next)
			}
		}
	}
	return head
}
