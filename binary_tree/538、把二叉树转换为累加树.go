package binary_tree

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	tAdd(root, &sum)
	return root
}

func tAdd(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	tAdd(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	tAdd(root.Left, sum)
}
