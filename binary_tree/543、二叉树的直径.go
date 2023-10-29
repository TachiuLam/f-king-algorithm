package binary_tree

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int
	diaMaxDepth(root, &maxDiameter)
	return maxDiameter
}

// 计算节点的最大直径: 返回该节点直径，并更新全局最大直径
func diaMaxDepth(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}
	leftDepth := diaMaxDepth(root.Left, maxDiameter)
	rightDepth := diaMaxDepth(root.Right, maxDiameter)
	diameter := leftDepth + rightDepth // 当前节点直径
	if diameter > *maxDiameter {
		*maxDiameter = diameter
	}
	return max(leftDepth, rightDepth) + 1 // 当前节点深度
}
