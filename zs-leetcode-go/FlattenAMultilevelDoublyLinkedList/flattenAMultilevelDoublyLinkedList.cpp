class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;
};

class Solution {
public:
    Node* flatten(Node* head) {
        if(head == nullptr) return head;
        aux(head->child, head);
        return head;
    }
    Node* aux(Node* child, Node* father) {
        if(child == nullptr) {
            if(father->next == nullptr) {
                return father;
            }
            return aux(father->next->child, father->next);
        }
        Node* father_next = father->next;
        father->next = child;
        child->prev = father;
        father->child = nullptr;
        Node* node = aux(child->child, child);
        if(father_next != nullptr) {
            father_next->prev = node;
            node->next = father_next;
            return aux(father_next->child, father_next);
        }
        return node;
    }
};
