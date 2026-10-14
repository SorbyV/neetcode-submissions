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
	if head.Next.Next == nil {
		node := head.Next
		node.Next = head
		head.Next = nil
		return node
	}
	n2 := head.Next
	n3 := n2.Next
	head.Next = nil
	for {
		if n3 == nil {
			break
		}
		n2.Next = head
		head = n2
		n2 = n3
		n3 = n3.Next
	}
	n2.Next = head
	return n2


}
