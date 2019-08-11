#include <iostream>
#include <algorithm>
#include <vector>
#include <map>

using namespace std;
class Node {
public:
    int val;
    vector<Node*> neighbors;

    Node() {}

    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};

class Solution {
public:
    map<int, Node**> nodeExists;
public:
    Node* cloneGraph(Node* node) {
        if (node == nullptr) {
            return nullptr;
        }
        Node* _begin = new Node(node->val);
        aux(&_begin, &node);
        return _begin;
    }
    auto exists(int val) {
        return [val](Node* v)-> bool {
            return v->val == val;
        };
    }
    void aux(Node** _node, Node** node) {
        for(auto iter =(*node)->neighbors.begin(); iter != (*node)->neighbors.end(); iter++) {
            int val = (*iter)->val;
            if(nodeExists.count(val) == 0){
                Node* _mp = new Node(val);
                _mp->neighbors.push_back(*_node);
                nodeExists[val] = &_mp;
                aux(&_mp, &(*iter));
            } else {
                (*nodeExists[val])->neighbors.push_back(*_node);
            }
        }
    }
};
int main()
{
    std::cout << "Hello world" << std::endl;
    return 0;
}

