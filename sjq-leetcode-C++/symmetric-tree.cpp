/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        if(root == NULL) return true;
        return isSame(root->left, root->right);
    }
    bool isSame(TreeNode* a, TreeNode*b){
        if(a == NULL) return !b;
        else if(b==NULL) return !a;
        else
            return (a->val==b->val) and isSame(a->left, b->right) and isSame(a->right, b->left);
    }
};
