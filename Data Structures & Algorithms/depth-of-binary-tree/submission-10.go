func find(root *TreeNode, curr int) int {
	if root == nil {
		fmt.Println("Reached an end, curr: ", curr)
		return 0
	}
	left := find(root.Left, curr+1)
	right := find(root.Right, curr+1)
	if left == 0 && right == 0 {
		return curr
	} else {
		return max(left, right)
	}
	return curr
}


func maxDepth(root *TreeNode) int {
    return(find(root, 1))
}
