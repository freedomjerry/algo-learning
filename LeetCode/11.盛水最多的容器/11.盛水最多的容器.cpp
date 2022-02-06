class Solution {
public:
    int maxArea(vector<int>& height) {
        int low = 0;
        int high = height.size() - 1;
        int maxCap = 0;
        while(low < high){
            if (height[low] < height[high]) {
                maxCap = max(maxCap, height[low] * (high - low));
                low++;
            }else{
                maxCap = max(maxCap, height[high] * (high - low));
                high--;
            }
        } 
        return maxCap;
    }
};