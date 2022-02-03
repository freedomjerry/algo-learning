class Solution {
public:
    void recFibonacci(int &k, int n1, int n2, int &count){
        int nFib = n1 + n2;
        if (k - nFib < 0){
            return;
        }else{
             recFibonacci(k, n2, nFib, count);
        }
        if (k - nFib >= 0){
            k = k - nFib;
            count += 1;
        }
        return;
    }
    int findMinFibonacciNumbers(int k) {
        int count = 0;
        recFibonacci(k, 1, 1, count);
        if(k == 1){
            count += 1;
        }
        return count;
    }
};