package binary_tree

func buildTreeFromMB(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootPostOrderIndex := len(postorder) - 1
	root := &TreeNode{Val: postorder[rootPostOrderIndex]}
	rootInorderIndex := 0
	for index, ino := range inorder {
		if ino == root.Val {
			rootInorderIndex = index
			break
		}
	}
	root.Left = buildTreeFromMB(inorder[:rootInorderIndex], postorder[:rootInorderIndex])
	root.Right = buildTreeFromMB(inorder[rootInorderIndex+1:], postorder[rootInorderIndex:rootPostOrderIndex])
	return root
}
