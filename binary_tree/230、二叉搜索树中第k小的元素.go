package binary_tree

func kthSmallest(root *TreeNode, k int) int {
	var res []int
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		res = append(res, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return res[k-1]
}

// 优化
func kthSmallestV2(root *TreeNode, k int) int {
	var (
		res  int
		rank int
	)
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		rank++
		if rank == k {
			res = node.Val
			return
		}
		traverse(node.Right)
	}
	traverse(root)
	return res
}
