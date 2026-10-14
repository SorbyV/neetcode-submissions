/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == head {
		return true
	}
    node1 := head
	node2 := head
	for {
		if node2 == nil || node2.Next == nil {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next.Next
		if node1 == node2 {
			return true
		}
	}
	return false
}
