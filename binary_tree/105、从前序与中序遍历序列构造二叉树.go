package binary_tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 前序遍历的第一个元素是根节点
	root := &TreeNode{Val: preorder[0]}
	rootInorderIndex := 0
	for index, ino := range inorder {
		if ino == root.Val {
			// 找到中序遍历中，root节点的index
			rootInorderIndex = index
			break
		}
	}
	// 构造左子树
	root.Left = buildTree(preorder[1:], inorder[:rootInorderIndex])
	// 构造右子树，根据 rootInorderIndex 在inorder的长度，计算出preorder 属于右子树的部分
	root.Right = buildTree(preorder[rootInorderIndex+1:], inorder[rootInorderIndex+1:])
	return root
}
