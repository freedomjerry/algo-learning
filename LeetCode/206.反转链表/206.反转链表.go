/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
if head == nil {
return nil
}
if head.Next == nil {
return head
}
slow := head
head = head.Next
slow.Next = nil
for head != nil{
fast := head.Next
head.Next = slow
slow = head
head = fast
}
return slow
}
