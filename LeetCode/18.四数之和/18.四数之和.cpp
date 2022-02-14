class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        vector<vector<int>> ans;
        if(nums.size() < 4){
            return ans;
        }
        sort(nums.begin(), nums.end());
       int n = nums.size();
        for(int i = 0; i < n - 3; i++ ){
            if ((long)nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] >target) {
                    break;
            }
            
            if(i > 0 && nums[i] == nums[i - 1]|| (long)nums[i] + nums[n - 3] + nums[n -2] + nums[n - 1] <  target){
                continue;
            }
            
            for(int j = i + 1; j < n - 2; j++){
                if ((long)nums[i] + nums[j] + nums[j + 1] + nums[j + 2] >target) {
                    break;
                }
                if(j > i + 1 && nums[j] == nums[j -1]|| (long)nums[i] + nums[j] + nums[n -2] + nums[n - 1] < target){
                    continue;
                }

                int low = j + 1;
                int high = n -1;
                while(low < high){
                    int sum = nums[i] + nums[j] + nums[low] + nums[high];
                    if(sum == target){
                        ans.push_back({nums[i], nums[j], nums[low], nums[high]});
                        while(low < high && nums[low] == nums[low + 1]) low++;
                        while(low < high && nums[high] == nums[high -1]) high--;
                        low++;
                        high--;
                    }else if (sum < target){
                        low++;
                    }else{
                        high--;
                    }
                }
            }
        }
        return ans;
    }
};