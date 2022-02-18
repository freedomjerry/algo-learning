func combinationSum(candidates []int, target int) (ans [][]int) {
    comb := []int{}
    var dfs func (target, index int)
    
    dfs = func(target, index int){
        if index == len(candidates){
            return
        }
        if target == 0 {
            ans = append(ans, append([]int(nil), comb...))
            return
        }
        dfs(target, index + 1)
        if target - candidates[index] >= 0 {
            comb = append(comb, candidates[index])
            dfs(target - candidates[index], index)
            comb = comb[:len(comb) -1]
        }
    }
    dfs(target, 0)
    return
    
}