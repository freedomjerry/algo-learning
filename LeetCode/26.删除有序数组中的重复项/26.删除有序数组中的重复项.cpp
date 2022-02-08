class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int n = nums.size();
        if(n == 0){
            return 0;
        }
        int now = 0;
        for(int i = 1; i < n ; i++){
            if(nums[i] == nums[now]){
                continue;
            }else{
                nums[now + 1] = nums[i];
                now++;
            }
        }
        return now + 1;
    }
};