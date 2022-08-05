package main

import "fmt"

func main()  {
	digit := "{}"
	fmt.Println()
}
//
//type combin []string
//
//func letterCombinations(digits string) []string {
//	char := map[int]string{}
//	for i:=2; i <= 6; i++ {
//		for j:=0; j<3; j++ {
//			char[i] += string('a' + (i-2)*3 + j)
//		}
//	}
//
//	char[7] = "pqrs"
//	char[8] = "tuv"
//	char[9] = "wxyz"
//	var combinations combin
//	combinations.getconbine(char, 0, digits, "")
//	return combinations
//}
//
//func (c *combin)getconbine(char map[int]string, layer int, digits string, comb string) {
//	if layer == len(digits){
//		if len(comb) != 0{
//			*c = append(*c, comb)
//		}
//		return
//	}
//	for _, v := range char[int(digits[layer]-'0')] {
//		c.getconbine(char, layer+1, digits, comb+string(v))
//	}
//	return
//}
