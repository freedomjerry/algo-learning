func threeSum(nums []int) (ans [][]int) {
    sort.Ints(nums)
    n := len(nums)
    for i:= 0; i < n - 2; i++ {
        if i > 0 && nums[i]==nums[i - 1] {
            continue
        }
        low := i + 1
        high := n - 1
        for low < high {
            sum := nums[i] + nums[low] + nums[high]
            if sum == 0 {
                temp := []int {nums[i], nums[low], nums[high]}
                ans = append(ans, temp)
                for low < high && nums[low] == nums[low + 1] {low++}
                for low < high && nums[high] == nums[high - 1] {high--}
                low ++
                high --
            }else if sum < 0 {
                low ++
            }else {
                high--
            }
        }
        
    }
    return
}