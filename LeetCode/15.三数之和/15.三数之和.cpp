#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        int size = nums.size();
        sort(nums.begin(), nums.end());
        vector<vector<int>> ans;
        for(int i = 0; i < size - 2; i++){
            if(nums[i] > 0){
                break;
            }
            if(i > 0 && nums[i] == nums[i - 1]){
                continue;
            }
            int low = i + 1;
            int high = size - 1;
            while(low < high){
                int sum = nums[high] + nums[low] + nums[i];
                if( sum == 0){
                    ans.push_back({nums[i], nums[low], nums[high]});
                    while(low < high && nums[low] == nums[low + 1]) low++;
                    while(low < high && nums[high] == nums[high - 1]) high--;
                    low++;
                    high--;
                }else if (sum > 0){
                    high--;
                }else {
                    low++;
                }
            }
        }
        return ans;
    }
    void PrintVector(vector<int> &v){
    for(vector<int>::iterator it = v.begin(); it != v.end(); it++)  {
        cout << *it << " ";
    }
    cout << endl;
}
};

int main(){
    vector<int> nums;
    int n;
    cin >> n;
    for (int i = 0 ; i < n; i++){
        int num;
        cin >> num;
        nums.push_back(num);
    }
    Solution sol;
    vector<vector<int>> ans = sol.threeSum(nums);
    for(vector<vector<int>>::iterator it = ans.begin(); it != ans.end(); it++){
        sol.PrintVector(*it);
    }
    return 0;
}