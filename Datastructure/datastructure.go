package main

import "fmt"

type BST struct {
	val int
	lchild, rchild *BST
}
func posArrayToBST(posArr []int) *BST {
 	return getBST(posArr, 0, len(posArr) - 1)
}

func getBST(arr []int, begin, end int) *BST {

	if begin > end {
		return nil
	}
	Node := &BST{}
	Node.val = arr[end]
	if begin == end {
		Node.lchild = nil
		Node.rchild = nil
		return Node
	}
	big := begin - 1
	left := begin
	right := end - 1
	for left < right {
		//if (arr[i] < arr[end]){
		//	big = i
		//}
		 mid := left + (right - left) >> 1
		 if arr[mid] < arr[end] {
		 	big = mid
		 	left = mid + 1
		 } else {
		 	right = mid - 1
		 }
	}
	Node.lchild = getBST(arr, begin, big)
	Node.rchild = getBST(arr, big + 1, end - 1)

	return Node
}

func main()  {
	arr := []int{1,3,2,6,8,7,5}
	BST := posArrayToBST(arr)
	fmt.Println(BST)
}