class Solution {
public:
    int binarySearch(vector<int>& nums, int target, int lower){
        int l = 0;
        int r = nums.size() - 1;
        int ans = -1;
        while(l <= r){
            int mid = (r - l) / 2 + l ;
            if(nums[mid] > target || (nums[mid] == target && lower == 1)){
                if(nums[mid] == target){
                    ans = mid;
                }
                r = mid -1;
            }else{
                if(nums[mid] == target){
                    ans = mid;
                }
                l = mid + 1;
            }
        }
        return ans;
    }
    vector<int> searchRange(vector<int>& nums, int target) {
        vector<int> ans;
        ans.push_back(binarySearch(nums, target, 1));
        ans.push_back(binarySearch(nums, target, 0));
        return ans;
    }
};