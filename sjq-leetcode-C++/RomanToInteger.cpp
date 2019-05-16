class Solution {
public:
    int romanToInt(string s) {
        map<char, int> m;
        m['I'] = 1; m['V'] = 5; m['X'] = 10; m['L'] = 50; m['C'] = 100; m['D'] = 500; m['M'] = 1000;
        int res = 0;
        for(int i = 0; i < s.size(); i++){
            if(s[i] == 'I'){
                if(s[i+1] == 'V'){
                    res += 4;
                    i++;
                    continue;
                }
                if(s[i+1] == 'X'){
                    res += 9;
                    i++;
                    continue;
                }
            }
            if(s[i] == 'X'){
                if(s[i+1] == 'L'){
                    res += 40;
                    i++;
                    continue;
                }
                if(s[i+1] == 'C'){
                    res += 90;
                    i++;
                    continue;
                }
            }
            if(s[i] == 'C'){
                if(s[i+1] == 'D'){
                    res += 400;
                    i++;
                    continue;
                }
                if(s[i+1] == 'M'){
                    res += 900;
                    i++;
                    continue;
                }
            }
            res += m[s[i]];
            cout << res;
        }
        return res;
    }
};
