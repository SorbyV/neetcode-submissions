/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func printList(head *ListNode) {
	node := head
	for {
		if node == nil {
			fmt.Println("")
			return 
		}
		fmt.Print("-", node.Val, "-")
		node = node.Next
	}

	
}
func length(head *ListNode) int {
	node := head
	counter := 0
	for {
		if node == nil {
			return counter
		}
		counter += 1
		node = node.Next
	}
	return counter
}

func Reverse(head *ListNode) *ListNode {
	
	if head.Next == nil {
		return head
	}
	node1 := head.Next
	if node1.Next == nil {
		node1.Next = head
		head.Next = nil
		return node1
	}
	node2 := node1.Next
	head.Next = nil
	for {
		// printList(node1)
		// printList(node2)
		if node2 == nil {
			node1.Next = head
			return node1
		}
		node1.Next = head
		head = node1
		node1 = node2
		node2 = node2.Next
	}
	node1.Next = head
	return node1
}
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
    counter := length(head)
	list2 := &ListNode{}
	list1 := head
	for i := 0; i < (counter+1)/2-1; i++ {
		list1 = list1.Next
	}
	list2.Val = list1.Next.Val
	list2.Next = list1.Next.Next
	list1.Next = nil
	list21 := Reverse(list2)
	// printList(list21)
	list11 := head
	
	list12 := list11.Next
	list22 := list21.Next
	for {
		// printList(head)
		if list12 == nil {
			if list21 == nil {
				return
			}
			list11.Next = list21
			return
		}
		list11.Next = list21
		list21.Next = list12
		list11 = list12
		list12 = list11.Next
		list21 = list22
		if list21 == nil {
			return
		}
		list22 = list22.Next
	}
}



