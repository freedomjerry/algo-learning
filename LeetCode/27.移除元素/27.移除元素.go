func removeElement(nums []int, val int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    now := 0
    for i:=0; i < n; i++{
        if nums[i] == val {
            continue;
        }else {
            nums[now] = nums[i]
            now++ 
        }
    }
    return now
}