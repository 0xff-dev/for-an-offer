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
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        if(nums.size() == 0) return NULL;
        return solve(nums, 0, nums.size()-1);
    }
    TreeNode* solve(vector<int>& nums, int left, int right){
        //从中间分
        if(left > right) return NULL;
        int mid = (left + right)/2;
        TreeNode* root = new TreeNode(nums[mid]);
        root->left = solve(nums, left, mid-1);
        root->right = solve(nums, mid+1, right);
        return root;
    }
};
