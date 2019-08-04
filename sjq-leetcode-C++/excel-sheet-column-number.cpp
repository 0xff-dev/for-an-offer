class Solution {
public:
    int titleToNumber(string s) {
        int res=0;
        int size = s.size();
        for(int i=0;i<size;i++){
            res += (s[i]-'A'+1) * pow(26,size-i-1);
        }
        return res;
    }
};
