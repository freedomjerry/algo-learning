class Solution {
public:
    int search(vector<int>& nums, int target) {
        int n = nums.size();
        if (!n) {
            return -1;
        }
        if(n == 1){
            return nums[0] == target? 0: -1;
        }
        int high = n - 1;
        int low = 0;
        while (low <= high){
            int mid = (low + high) / 2;
            if(nums[mid] == target){
                return mid;
            }else if (nums[mid] < nums[0]){
                if(nums[mid] < target && target <= nums[n - 1] ){
                    low = mid + 1;
                }else{
                    high = mid -1;
                }
            }else if (nums[mid] >= nums[0]){
                if(nums[mid] > target && nums[0] <= target){
                    high = mid - 1;
                }else{
                    low = mid + 1;
                }
            }
        }
        return -1;
    }
};