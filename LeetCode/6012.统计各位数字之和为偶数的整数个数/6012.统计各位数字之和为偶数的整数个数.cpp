#include<iostream>
using namespace std;
class Solution {
public:
    int countEven(int num) {
        int count = 0;
        int sum = 0;
        int temp = num/10;
        while(temp != 0){
            sum += temp % 10;
            temp /= 10;
        }
        cout << sum << endl;      
        count += (num % 10)/2;
        cout << count << endl;
        count += (sum % 2 == 0 ? (num / 10) * 5 :(num / 10) * 5 - 1);
        if((sum + (num % 10))%2 == 0 && sum % 2 == 1){
            count++;
        }
        return count;
    }
};
int main(){
    int num = 63;
    Solution s;
    cout << s.countEven(num) << endl;
}