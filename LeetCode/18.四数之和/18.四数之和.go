func fourSum(nums []int, target int) (ans [][]int) {
    n := len(nums)
    sort.Ints(nums)
    for i:= 0; i < n - 3; i++ {
        if i > 0 && nums[i] == nums[i -1] {
            continue
        }
        for j := i + 1; j < n -2; j++ {
            if j > i + 1 && nums[j] == nums[j -1]{
                continue
            }
            low := j + 1
            high := n -1
            for low < high {
                sum := nums[i] + nums[j] + nums[low] + nums[high]
                if sum == target {
                    ans = append(ans, []int{nums[i], nums[j], nums[low], nums[high]})
                    for low < high && nums[low] == nums[low + 1] {low++}
                    for low < high && nums[high] == nums[high - 1]{high--}
                    low ++
                    high -- 
                }else if sum < target{
                    low++
                }else {
                    high--
                }
            }
        }
    }
    return
}