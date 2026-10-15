/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isEdgeNode(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := root.Left
	root.Left = root.Right
	root.Right = node
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
