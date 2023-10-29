package binary_tree

func inorderTraversal(root *TreeNode) []int {
	var (
		res     []int
		inorder func(root *TreeNode)
	)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return res
}
