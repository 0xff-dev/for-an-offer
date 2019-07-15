class Solution {
public:
    int mySqrt(int x) {
        if(x == 0) return 0;
        double k = 1.0;
        while(abs(k*k - x) > 1e-4){
            k = (k+x/k)/2.0;
        }
        return int(k);
    }
};
