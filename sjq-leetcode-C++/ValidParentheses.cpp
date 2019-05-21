class Solution {
public:
    bool isValid(string s) {
        if(s.size() == 0) return true;
        if(s.size() == 1) return false;
        vector<char> vec;
        vec.push_back('}');
        for(int i = 0; i < s.size(); i++){
            if(s[i] == '(' ||s[i] == '{' || s[i] =='[' )
                vec.push_back(s[i]);
            else{
                if((s[i] == ')'&& vec.back() == '(')||(s[i] == '}' && vec.back() == '{')||(s[i] == ']' && vec.back() == '[') ){
                    vec.pop_back();
                }
                else return false;
            }
        }
        if(vec.size() == 1)
            return true;
        return false;
    }
};
