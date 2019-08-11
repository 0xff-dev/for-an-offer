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
    map<Node*, Node*> _map;
public:
    Node* cloneGraph(Node* node) {
        if (node == nullptr) {
            return nullptr;
        }
        if(_map.count(node) == 0) {
            _map[node] = new Node(node->val);
            for(auto iter : node->neighbors) {
                Node* nodes = cloneGraph(iter);
                _map[node]->neighbors.push_back(nodes);
            }
        }
        return _map[node];
    }
};
