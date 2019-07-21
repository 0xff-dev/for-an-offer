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
    bool flag = true;
    bool isBalanced(TreeNode* root) {
        if(root == NULL) return true;
        int left = depth(root);
        return flag;
    }
    int depth(TreeNode* root){
        if(root == NULL) return 0;
        int left = depth(root->left);
        int right = depth(root->right);
        if(abs(left-right)>1) flag=false;
        return max(left, right)+1;
    }
};
