class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int l = 0;
        int r = nums.size() - 1;
        int mid;
        while (l <= r){
            mid = (r - l) /2 + l;
            if(nums[mid] == target){
                break;
            }else if(nums[mid] < target){
                l = mid + 1;
            }else {
                r = mid - 1;
            }
        }
        if(nums[mid] < target){
            mid ++;
        }
        return mid;
    }
};