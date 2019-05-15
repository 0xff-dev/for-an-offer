class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int length = s.size();
        int sum = 0;
        int begin = 0;
        if(length == 1)
            return 1;
        
        int i = 0;
        for(int j = 0; j < length; j++){
            for(int k = i; k < j; k++){
                if(s[k] == s[j]){
                    i = k+1;
                    break;
                }
            }
            if(j-i+1 > sum)
                sum = j-i+1;
        }
        return sum;
    }
};
