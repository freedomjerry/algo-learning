func abs(a int, b int) int{
    if a - b < 0 {
        return b - a
    }else {
        return a - b
    }
}
func threeSumClosest(nums []int, target int) (ans int) {
    sort.Ints(nums)
    ans = nums[0] + nums[1] + nums[2]
    n := len(nums)
    if ans == target || n == 3{
        return
    }   
    for i := 0; i < n - 2; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        low := i + 1
        high := n -1
        for low < high {
            sum := nums[i] + nums[low] + nums[high]
            if abs(target, sum) < abs(target, ans) {
                ans = sum
            } 
            if sum == target {
                return 
            }else if sum < target {
                low ++
            }else {
                high--
            }
        }
    }
    return
}