class Solution {
public:
    int getSize(vector<string>& strs){
        int size = 10000;
        for(int i = 0; i < strs.size(); i++){
            if(size > strs[i].size())
                size = strs[i].size();
        }
        return size;
    }
    string longestCommonPrefix(vector<string>& strs) {
        if(strs.size() == 0) return "";
        if(strs.size() == 1) return strs[0];
        sort(strs.begin(), strs.end());
        int j;
        for(int i = 0; i< getSize(strs); i++){
            bool flag = true;
            for(j = 0; j < strs.size()-1; j++){
                if(strs[j][i] != strs[j+1][i]){
                    flag = false;
                    break;
                }
            }
            if(!flag)
                return strs[0].substr(0, i);
        }
        return strs[0];
    }
};
