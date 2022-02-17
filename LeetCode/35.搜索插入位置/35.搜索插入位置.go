func searchInsert(nums []int, target int) int {
    i := sort.SearchInts(nums, target)
    return i
}