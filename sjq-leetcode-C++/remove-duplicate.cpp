/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if(head == NULL) return NULL;
        ListNode* node = head;
        while(node->next){
            if(node->val == node->next->val){
                if(node->next->next !=NULL)
                    node->next = node->next->next;
                else node->next = NULL;
            }
            else
                node = node->next;
        }
        return head;
    }
};
