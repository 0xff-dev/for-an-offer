class Solution {
public:
    int majorityElement(vector<int>& nums) {
        map<int, int> mapNums;
        int res=0;
        for(int i=0; i< nums.size(); i++){
            mapNums[nums[i]]++;
            if(mapNums[nums[i]]>nums.size()/2)
                res=nums[i];
        }
        return res;
    }
};
