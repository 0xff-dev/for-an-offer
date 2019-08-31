class Solution {
public:
    int myAtoi(string str) {
        int i=0;
        int flag = 1;
        long long int res=0;
        while(i < str.size() && str[i] == ' ') i++;
        if(i == str.size())
            return 0;
        if(str[i] == '-') {
            flag=-1;
            i++;
        }
        else if(str[i] == '+') {
            flag=1;
            i++;
        }
        while(i<str.size() && str[i]>='0' && str[i]<='9'){
            res = res*10 +str[i++]-'0';
            if(res >= 2147483648){
                if(flag == 1) return INT_MAX;
                return INT_MIN;
            }
        }
        return res*flag;
        
    }
};
