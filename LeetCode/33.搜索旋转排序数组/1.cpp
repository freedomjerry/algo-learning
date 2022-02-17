#include<iostream>
#include<vector>
using namespace std;
int main(){
    vector<int> nums = {2, 2};
    int target = 2;
    int n = nums.size();
    int l = 0;
    int r = 0;
    while(nums[l - 1] == target && l != 0) l--;
    while(nums[r + 1] == target && r != n -1) r++;
    cout  << l << " " << r << endl;
}
