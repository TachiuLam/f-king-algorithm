package binary_tree

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 如果只存在右子树
		if root.Left == nil {
			return root.Right
		}
		// 如果只存在左子树
		if root.Right == nil {
			return root.Left
		}
		// 如果存在左右子树 , 右子树最小的节点须提出来，替换要删除的根节点
		minNode := getMin(root.Right)
		root.Right = deleteNode(root.Right, minNode.Val)
		root.Val = minNode.Val // 替换节点
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

// 找到右子树的最小节点
func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

//leetcode submit region end(Prohibit modification and deletion)
