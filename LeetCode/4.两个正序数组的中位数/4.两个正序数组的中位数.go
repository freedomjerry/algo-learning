package main

import (
	"fmt"
	"sort"
)

//
//
//func min(x int, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//func getmedianNumber(nums1 []int, nums2 []int, k int) int {
//	len1 := len(nums1)
//	len2 := len(nums2)
//	p1 := 0
//	p2 := 0
//	for  {
//		if p1 == len1 {
//			return nums2[p2 + k - 1]
//		}else if p2 == len2 {
//			return nums1[p1 + k - 1]
//		}else if k == 1 {
//			return min(nums1[p1], nums2[p2])
//		}
//
//		q1 := min(p1 + k / 2 - 1, len1 - 1)
//		q2 := min(p2 + k / 2 - 1, len2 - 1)
//		if nums1[q1] <= nums2[q2] {
//			k = k - q1 + p1 - 1
//			p1 = q1 + 1
//		}else {
//			k = k - q2 + p2 - 1
//			p2 = q2 + 1
//		}
//	}
//}
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	len1 := len(nums1)
//	len2 := len(nums2)
//	total := len1 + len2
//	if total % 2 == 1 {
//		return float64(getmedianNumber(nums1, nums2, (total + 1) / 2))
//	}else {
//		return float64(getmedianNumber(nums1, nums2, total / 2) + getmedianNumber(nums1, nums2, total / 2 + 1)) / 2
//	}
//}
func main()  {
	a := []int {5, 3, 3, 4, 2, 6}
	sort.Ints(a)
	fmt.Println(a)
}
