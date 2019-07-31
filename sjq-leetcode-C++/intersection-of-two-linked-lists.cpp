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
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if(headA==NULL or headB==NULL) return NULL;
        set<ListNode*> setA;
        while(headA){
            setA.insert(headA);
            headA = headA->next;
        }
        while(headB){
            if(setA.find(headB) != setA.end())
                break;
            headB = headB->next;
        }
        return headB;
    }
};
