package binary_tree

// Node Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse116(root.Left, root.Right)
	return root
}

// 三叉树遍历
func traverse116(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	// 连接左右节点
	node1.Next = node2

	// 连接左子树的左右节点
	traverse116(node1.Left, node1.Right)
	// 连接右子树的左右节点
	traverse116(node2.Left, node2.Right)
	// 连接跨越父节点的两个子节点
	traverse116(node1.Right, node2.Left)
}
