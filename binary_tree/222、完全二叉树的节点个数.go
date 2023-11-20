package binary_tree

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		lh, rh int
		l, r   = root, root
	)
	for l != nil {
		l = l.Left
		lh++
	}
	for r != nil {
		r = r.Right
		rh++
	}
	if lh == rh { // 满二叉树
		return int(math.Pow(float64(2), float64(lh))) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
