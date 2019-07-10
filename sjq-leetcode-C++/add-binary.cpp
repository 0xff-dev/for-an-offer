class Solution {
public:
    string addBinary(string a, string b) {
        //
        int i =a.size() -1, j =b.size() -1;
        string c;
        int temp = 0;
        while(i >=0 && j >=0){
            int aa = a[i] - '0', bb = b[j] - '0';
            int plus = aa +bb +temp;
            if(plus == 2){
                temp = 1;
                c.insert(0, "0");
            }else if(plus == 3){
                temp = 1;
                c.insert(0, "1");
            }else if(plus == 1){
                temp = 0;
                c.insert(0, "1");
            }else{
                temp = 0;
                c.insert(0, "0");
            }
            i--;
            j--;
        }
        while( i >= 0){
            int aa = a[i] - '0';
            int plus = aa +temp;
            if(plus == 2){
                temp = 1;
                c.insert(0, "0");
            }else if(plus == 1){
                temp = 0;
                c.insert(0, "1");
            }else{
                temp = 0;
                c.insert(0, "0");
            }
            i--;           
        }
        while( j >= 0){
            int bb = b[j] - '0';
            int plus = bb +temp;
            if(plus == 2){
                temp = 1;
                c.insert(0, "0");
            }else if(plus == 1){
                temp = 0;
                c.insert(0, "1");
            }else{
                temp = 0;
                c.insert(0, "0");
            }
            j--;  
        }
        if(temp)
            c.insert(0, "1");
        return c;
    }
};
