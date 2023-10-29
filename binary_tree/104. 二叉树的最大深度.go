package binary_tree

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var depth = 0

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	depth = max(leftDepth, rightDepth) + 1
	return depth
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
