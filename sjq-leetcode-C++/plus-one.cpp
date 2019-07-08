class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        vector<int> res = digits;
        int size = digits.size();
        int i=1;
        res[size -1] += 1;
        if(res[i=size -1] == 10){
            for(i = size-1; i>0; i--){
                res[i] = 0;
                res[i-1] += 1;
                if(res[i -1] != 10)
                    break;
            }
        }
        if(i == 0){
            if(res[i] == 10){
                res[i] = 0;
                res.insert(res.begin(), 1);
            }
        }
        return res;
    }
};
