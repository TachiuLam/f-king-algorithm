package binary_tree

func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	traverse(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
