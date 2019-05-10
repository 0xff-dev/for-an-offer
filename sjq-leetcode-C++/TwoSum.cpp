#include <algorithm>
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> res;
        for(int i = 0; i<nums.size(); i++){
            int temp = target - nums[i];
            if(find(nums.begin(), nums.end(), temp) != nums.end()){
                for(int j = i+1; j < nums.size(); j++){
                    if(temp == nums[j]){
                        res.push_back(i);
                        res.push_back(j);
                    }
                }
            }
        }
        
        return res;
    }
};
