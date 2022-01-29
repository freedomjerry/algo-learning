class Solution {
public:
    int getmiddleNums(const vector<int>& nums1,const vector<int>& nums2, int k){
        int len1 = nums1.size();
        int len2 = nums2.size();
        int p1 = 0;
        int p2 = 0;
        while(true){
            if (p1 == len1) {
                return nums2[k + p2 - 1];
            }else if(p2 == len2){
                return nums1[k + p1 - 1];
            }else if(k == 1){
                return min(nums1[p1], nums2[p2]);
            }
            int q1 = min(p1 + k / 2 - 1 , len1 - 1);
            int q2 = min(p2 + k /2 -1, len2 - 1);
            int num1 = nums1[q1];
            int num2 = nums2[q2];
            if (num1 <= num2) {
                k -= q1 - p1 + 1;
                p1 = q1 + 1;
            }else{
                k -= q2 - p2 + 1;
                p2 = q2 + 1;
            }
        }
    }
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int k = nums1.size() + nums2.size();
        if( k % 2 != 0){
            return getmiddleNums(nums1, nums2, (k + 1)/ 2 );
        }else{
            return (getmiddleNums(nums1, nums2, k / 2) + getmiddleNums(nums1, nums2, k / 2 +1)) * 1.0 / 2; 
        }
    }
};