package binary_tree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 拉平左右子树
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	// 将左子树置为nil， 右子树为原先左子树
	root.Right = left
	root.Left = nil

	// 将原右子树 接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}
