class Solution {
public:
    string convertToTitle(int n) {
        string res;
        while(n>0){
            n--;
            int temp = n%26;
            char s = 'A'+temp;
            res =s+res;
            n /= 26;
        }
        return res;
    }
};
