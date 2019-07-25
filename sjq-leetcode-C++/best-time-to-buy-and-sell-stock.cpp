class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if(prices.size() == 0) return 0;
        while(prices.back() ==0) prices.pop_back();
        int i=0;
        int res=0;
        while(i<prices.size()-1){
            int j = i+1;
            while(j<prices.size()){
                res = max(res, prices[j]-prices[i]);
                j++;
            }
            i++;
        }
        return res;
    }
};
