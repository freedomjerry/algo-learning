package main

import "fmt"

func firstMissingPositive(nums []int) int {
    var swap func(a int, b int)
    swap = func(a int,b int){
        c := nums[a];
        nums[a] = nums[b];
        nums[b] = c;
    }
    for i:= 0; i < len(nums); i++{
        if nums[i] > 0 && nums[i] != i+1{
            if nums[i] <= len(nums) && nums[nums[i] - 1] != nums[i]{
                swap(i, nums[i] -1)
                i--;
            } else {
                nums[i] = 0;
            }
            
        }
    }
    num := 0
    for num < len(nums)&& nums[num] == num + 1  { num++ }
    return num + 1
}
func main()  {
    nums := []int{2, 1}
    fmt.Println(firstMissingPositive(nums))
}