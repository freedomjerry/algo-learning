func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    now := 0
    for i:=0; i < n; i++{
        if nums[now] != nums[i] {
            nums[now + 1] = nums[i]
            now++
        }
    }
    return now + 1
}