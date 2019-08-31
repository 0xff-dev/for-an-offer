class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        // 调转三次
        if(nums.size()<2) return;
        k %= nums.size();
        reverse(nums, 0, nums.size()-1);
        reverse(nums, 0, k-1);
        reverse(nums, k, nums.size()-1);
    }
    
    void reverse(vector<int>& nums, int l, int r){
        while(l < r)
            // 交换前后两个元素
            swap(nums[l++], nums[r--]);
    }
};
