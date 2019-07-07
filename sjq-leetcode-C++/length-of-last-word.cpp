class Solution {
public:
    int lengthOfLastWord(string s) {
        int len = s.size();
        if(len < 1)
            return 0;
        int i = len -1;
        int res = 0;
        for(i; i >= 0; i--){
            if(s[i] != ' ')
                res++;
            else if(res > 0)
                break;
        }
        return res;
    }
};
