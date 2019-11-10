#include<iostream>

class Node {
public:
    int val;
    Node* next;
    Node* random;

    Node() {}

    Node(int _val, Node* _next, Node* _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};

class Solution {
public:
    Node* copyRandomList(Node* head) {
        // copy list
        if(head == nullptr) {
            return nullptr;
        }
        Node* walker, *aux_poniter, *res_head;
        walker = head;
        while(walker != nullptr) {
            Node* tmp = new Node(walker->val, walker->next, nullptr);
            walker->next = tmp;
            walker = tmp->next;
        }
        // copy random
        walker = head;
        while(walker != nullptr) {
            if(walker->random != nullptr) {
                walker->next->random = walker->random->next;
            }
            walker = walker->next->next;
        }

        // get aim list
        walker = head, aux_poniter = head->next, res_head = head->next;
        while(walker != nullptr) {
            walker->next = aux_poniter->next;
            walker = aux_poniter->next;
            if(walker != nullptr) {
                aux_poniter->next = walker->next;
                aux_poniter = walker->next;
            }
        }
        return res_head;
    }
};

