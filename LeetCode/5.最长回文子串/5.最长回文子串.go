package  main

import "fmt"

func longestPalindrome(s string)(ans string){
	n := len(s)
	if n < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	dp := make([][]int, n)
	for i:= 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i:= 0; i< n ; i++{
		dp[i][i] = 1
	}
	//fmt.Println(dp)
	for L := 2; L <= n; L++ {
		for j:=0; j<n; j++{
			r:= j + L -1
			if r >= n {
				break;
			}
			if s[j] != s[r] {
				dp[j][r] = 0
			}else{
				if L < 4 {
					dp[j][r] = 1
				}else{
					dp[j][r] = dp[j+1][r-1]
				}
			}

			if dp[j][r] == 1 && L > maxLen {
				maxLen = L
				begin = j
			}
		}
	}
	return s[begin: begin + maxLen]
}

func main()  {
	var s string
	fmt.Scan(&s)
	fmt.Println(longestPalindrome(s))
}