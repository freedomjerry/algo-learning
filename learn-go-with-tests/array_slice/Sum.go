package array_slice

//func Sum(nums [5]int) int  {
//	sum := 0
//	//for i := 0; i < len(nums); i++ {
//	//	sum += nums[i]
//	//}
//	for _, number := range nums {
//		sum += number
//	}
//	return sum
//}
func Sum(nums []int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum

}

func SumAll(numbersToSum ...[]int) (sums []int) {
	//lenthOfNumbers := len(numbersToSum)
	//sums = make([]int, lenthOfNumbers)
	for _, numbers := range numbersToSum{
		sums = append(sums, Sum(numbers))
	}
	return
}
func SumAllTails(numbersToTail ...[]int) (sums []int) {

	for _, number := range numbersToTail {
		if len(number) == 0 {
			sums = append(sums, 0)
		}else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}