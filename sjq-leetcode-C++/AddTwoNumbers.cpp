/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
#include<stdlib.h>
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) 
    {
        ListNode* r=new ListNode(NULL);
        ListNode* c=r;
        int temp=0;
        while(l1 || l2)
        {
            if(l1) {
                temp += l1->val;
                l1 = l1->next;
            }
            if(l2) {
                temp += l2->val;
                l2 = l2->next;
            }
            c->next=new ListNode(temp%10);
            temp/=10;
            c=c->next;
        }
        if(temp)
            c->next=new ListNode(1);
        return r->next;
    }
};
