class Solution {
public:
    string countAndSay(int n) {
        if(n == 1) return "1";
        string res = "";
        string s = "";
        else{
            s = countAndSay(n-1);
            for(int i = 0; i< s.size(); i++){
                int j = 0;
                while(++j){
                    if((i+j) >= s.size()) break; //超出字符串长度
                    if(s[i] != s[i+j]) break; //不一致
                }
                res += to_string(j);
                res += s[i];
                i = i+j -1;
            }
            
        }
        return res;
    }
};
