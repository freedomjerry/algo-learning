class Solution {
public:
    void dfs(vector<int>& candidates, vector<vector<int>> &ans, vector<int> &comb,int target, int index){
        int n = candidates.size();
        if(index >= n){
            return;
        }
        if(target == 0){
            ans.emplace_back(comb);
            return;
        }
        dfs(candidates, ans, comb, target, index + 1);
        if(target - candidates[index] >= 0){
            comb.emplace_back(candidates[index]);
            dfs(candidates, ans, comb, target - candidates[index], index);
            comb.pop_back();
        }

    }
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        vector<vector<int>> ans;
        vector<int> comb;
        dfs(candidates, ans, comb, target, 0);
        return ans;
    }
};