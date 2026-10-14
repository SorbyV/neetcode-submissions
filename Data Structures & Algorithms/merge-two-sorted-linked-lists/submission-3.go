/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func printList(head *ListNode) {
// 	node := head
// 	for {
// 		if node == nil {
// 			fmt.Println(" ")
// 			return
// 		}
// 		fmt.Print("-", node.Val, "-")
// 		node = node.Next
// 	}

// }
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		0,
		nil,
	}
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}

	node := head

    for {
		// printList(head)
		if list1 == nil && list2 != nil {
			// fmt.Println("List1 finished: ", node.Val, list2.Val)
			node.Val = list2.Val
			node.Next =list2.Next

			// printList(head)
			return head
		}
		if list2 == nil && list1 != nil {
			// fmt.Println("List2 finished: ",  node.Val, list1.Val)
			node.Val = list1.Val
			node.Next =list1.Next

			// printList(head)
			return head
		}
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				// fmt.Println("List1 has smaller value: ",  node.Val, list1.Val, list2.Val)
				node.Val = list1.Val
				list1 = list1.Next
			} else {
				// fmt.Println("List2 has smaller value: ", node.Val, list1.Val, list2.Val)

				node.Val = list2.Val
				list2 = list2.Next
			}
			node.Next = &ListNode{}
			node = node.Next
		}
		
	}
	return head
}
