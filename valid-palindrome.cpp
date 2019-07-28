class Solution {
public:
    bool isPalindrome(string s) {
        transform(s.begin(), s.end(), s.begin(),  ::tolower);
        int left=0, right=s.size()-1;
        while(left<right){
            while(!(isalpha(s[left]) or isalnum(s[left])) and left<right) left++;
            while(!(isalpha(s[right]) or isalnum(s[right])) and left<right) right--;
            if(left > right) break;
            if(s[left] != s[right])
                return false;
            left++;
            right--;
        }
        return true;
    }
};
