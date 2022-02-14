class Solution {
public:
    void swapVector(vector<int>& nums, int i, int j){
        int k;
        k = nums[i];
        nums[i] = nums[j];
        nums[j] = k;
    }
    void nextPermutation(vector<int>& nums) {
        int n = nums.size();
        int slow = n -1;
        int quick = n -2;
        int tag = 0;
        for(quick; quick >= 0; quick--){
            if(tag == 1){
                break;
            }
            while(slow > quick){
                if(nums[slow] > nums[quick]){
                    swapVector(nums, slow, quick);
                    tag = 1;
                    sort(nums.begin()+ quick + 1, nums.end());
                    break;
                }else{
                    slow--;
                }
            }
            slow = n - 1;
        }
        if(tag == 1){
            return;
        }else{
            sort(nums.begin(), nums.end());
            return;
        }
    }
};