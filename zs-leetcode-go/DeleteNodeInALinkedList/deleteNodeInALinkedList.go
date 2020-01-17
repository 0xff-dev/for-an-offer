package DeleteNodeInALinkedList

import "for-an-offer/zs-leetcode-go/types"
// only give a node, does not have head. 
func deleteNode(node *types.ListNode) {
    *node = *(node.Next)
}
