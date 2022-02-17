func binarySearch(nums []int, target int, lower int) (ans int){
    ans = -1
    l := 0
    r := len(nums) - 1;
    for l <= r {
        mid := (r - l ) / 2 + l
        if nums[mid] > target || (nums[mid] == target && lower == 1){
            if nums[mid] == target {
                ans = mid
            }
            r = mid - 1
        }else {
            if nums[mid] == target {
                ans = mid
            }
            l = mid + 1
        }
    }
    return
}
func searchRange(nums []int, target int) []int {
    left := binarySearch(nums, target, 1)
    right := binarySearch(nums, target, 0)
    return []int {left, right}
}


// 2
func searchRange(nums []int, target int) []int {
    left := sort.SearchInts(nums, target)
    if left == len(nums) || nums[left] != target {
        return []int {-1, -1}
    }
    right := sort.SearchInts(nums, target + 1) - 1;
    return []int {left, right}
}