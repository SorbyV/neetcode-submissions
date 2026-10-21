/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(root *TreeNode, n int, balance bool) (int, bool) {
	if root == nil {
		return n, balance
	}
	left, bleft := height(root.Left, n, balance)
	right, bright := height(root.Right, n, balance)
	if left - right < -1 || left - right > 1 {
		fmt.Println("disbalanced: ", left, right)
		balance = false
	}
	balance = balance && (bleft&&bright)
	return max(left+1, right+1), balance
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	balance := true
    _, balance = height(root, 1, balance)
	return balance
}
