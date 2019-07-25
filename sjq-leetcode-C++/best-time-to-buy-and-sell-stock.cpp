class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if(prices.size() == 0) return 0;
        while(prices.back() ==0) prices.pop_back();
        int i=0;
        int res = 0, temp = prices[0];
        while(i<prices.size()){
            res = max(res, prices[i]-temp);
            temp = min(temp, prices[i]);
            i++;
        }
        return res;
    }
};
