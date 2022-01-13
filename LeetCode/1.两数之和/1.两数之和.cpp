#include<iostream>
#include<vector>
#include<unordered_map>
#include <algorithm>
using namespace std;

class Solution{
    public:
        vector<int> twoSum(vector<int>& nums, int target){
            unordered_map<int, int> hashtable; //unorder_map 用hash实现的无序map， key , value
            for(int i = 0; i < nums.size(); i++){ //O(n)循环，对每一个数判断与目标值的差是否在hash表中
                auto it = hashtable.find(target - nums[i]);//找与target的差
                if (it != hashtable.end()){
                    return {it -> second, i};
                }
                hashtable[nums[i]] = i;
            }
            return {};
        }
};
void MyPrint(int val){
    cout << val << " ";
}
int main (){
    vector<int> nums;
    for(int i = 0; i < 100; i++){
        nums.push_back(i);
    }
    int target = 104;
    Solution solution;
    auto result = solution.twoSum(nums, target);
    cout << result.size() << endl;
    for_each(result.begin(), result.end(), MyPrint);
    return 0;
}