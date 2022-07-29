package main

import "fmt"

func isMatch(s string, p string) bool {
	l1 := len(s)
	l2 := len(p)
	dp := make([][]bool, l1 + 1)
	for i:=0; i<= l1; i++ {
		dp[i] = make([]bool, l2 + 1)
	}
	match := func (i int, j int)bool{
		if i == 0 {
			return false
		}else if p[j - 1] == '.' {
			return true
		}
		return s[i - 1] == p[j - 1]
	}

	dp[0][0] = true
	for i := 0; i <= l1; i++ {
		for j:=1; j <= l2; j++{
			if p[j - 1] == '*'{
				dp[i][j] = dp[i][j - 2]
				if match(i, j - 1){
					dp[i][j] = dp[i][j] || dp[i - 1][j]
				}
			}else {
				if match(i,j){
					dp[i][j] = dp[i - 1][j - 1]
				}
			}
		}
	}
	return dp[l1][l2]
}

func main()  {
	var s string
	var p string
	fmt.Scan(&s, &p)
	fmt.Println(isMatch(s, p))
}