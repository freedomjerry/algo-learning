#include<iostream>
#include<vector>
#include<algorithm>
#include<unordered_map>
using namespace std;
class Solution {
public:
    void getIndex(int *index, unordered_map<int ,int> hash){
        if(hash[index[0]] < hash[index[1]]){
                    swap(index[0], index[1]);
        }
        for(auto &elem : hash){
            if(elem.first == index[0] || elem.first == index[1]){
                continue;
            }
            if(elem.second > hash[index[1]]){
                index[1] = elem.first;
                if(hash[index[0]] < hash[index[1]]){
                    swap(index[0], index[1]);
                }
            }
        }
    }
    int minimumOperations(vector<int>& nums) {
        unordered_map<int ,int> odds;
        unordered_map<int, int> evens;
        int index[2];
        int index1[2];
        int time =0;
        for(int i = 0; i < nums.size(); i+=2){ 
            auto it = odds.find(nums[i]);
            if (it != odds.end()){
                it->second++;
            }else{
                odds[nums[i]] = 1;
                if(time < 2){
                    index[time++] = nums[i];
                }
            }
            
        }
        time = 0;
        for(int i = 1; i < nums.size(); i+=2){ 
            auto it = evens.find(nums[i]);
            if (it != evens.end()){
                it->second++;
            }else{
                evens[nums[i]] = 1;
                if(time < 2){
                    index1[time++] = nums[i];
                }
            }
        }
        cout << index[0]<<" "<<index[1] << endl;
        getIndex(index, odds);
        cout << index[0]<<" "<<index[1] << endl;
        cout << odds[index[0]]<<" "<<odds[index[1]] << endl;
        getIndex(index1, evens);
        cout << index1[0]<<" "<<index1[1] << endl;
        cout << evens[index1[0]]<<" "<<evens[index1[1]] << endl;
        int maxNum = 0;
        for(int i = 0; i < 2; i++){
            for(int j = 0; j < 2; j++){
                if(index[i] != index1[j]){
                    maxNum = max(maxNum, odds[index[i]] + evens[index1[j]]);
                }
            }
        }
        return nums.size() - maxNum;
    }        
};
int main(){
    vector<int> nums = {1,2,2,2,2};
    Solution s;
    int ans = s.minimumOperations(nums);
    cout << nums.size() << endl;
    cout << ans;
    return 0;
}