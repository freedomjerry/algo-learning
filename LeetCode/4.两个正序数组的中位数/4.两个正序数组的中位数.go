
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
//func main()  {
//	a := []int {5, 3, 3, 4, 2, 6}
//	sort.Ints(a)
//	fmt.Println(a)
//}
package main

import "fmt"

func findMedianSortArrays(nums1 []int, nums2 []int)(ans float64){
	totallen := len(nums1) + len(nums2)
	mid := totallen / 2
	if totallen % 2 == 1 {
		return float64(findKth(nums1, nums2, mid + 1))
	}else {
		return float64((findKth(nums1, nums2, mid) + findKth(nums1, nums2, mid + 1)) / 2)
	}
	return 0
}

func findKth(nums1 []int, nums2 []int, k int)(ans int)  {
	r1 := len(nums1)
	r2 := len(nums2)
	l1, l2 := 0, 0
	for {
		if l1 == r1 {
			return nums2[l2 + k - 1]
		}
		if l2 == r2 {
			return nums1[l1 + k - 1]
		}
		if k == 1{
			return min(nums1[l1], nums2[l2])
		}
		mid := k / 2
		newl1 := min(l1 + mid, r1) - 1
		newl2 := min(l2 + mid, r2) - 1
		if nums1[newl1] <= nums2[newl2]{
			k -= (newl1 - l1 + 1)
			l1 = newl1 + 1
		}else{
			k -= (newl2 - l2 + 1)
			l2 = newl2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return  x

	}else {
		return  y
	}
}

func max(x, y int) int {
	if x > y {
		return x

	}else {
		return y
	}
}

func main()  {
	var n int
	var m int
	fmt.Scan(&n , &m)
	nums1 := make([]int, n)
	nums2 := make([]int, m)
	for i:= 0; i < n; i++{
		fmt.Scan(&nums1[i])
	}
	for j:= 0; j < m; j++ {
		fmt.Scan(&nums2[j])
	}
	ans := findMedianSortArrays(nums1, nums2)
	fmt.Println(ans)
}