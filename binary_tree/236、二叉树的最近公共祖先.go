package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return find(root, p.Val, q.Val)
}

func find(root *TreeNode, v1, v2 int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == v1 || root.Val == v2 { // root值等于q或p，且树中一定存在p、q，则该节点必定是最近公共祖先
		return root
	}

	left := find(root.Left, v1, v2)
	right := find(root.Right, v1, v2)
	// 如果左右子树都不为空，说明该节点是q、p的最近公共祖先
	if left != nil && right != nil {
		return root
	}

	// 左子树查找不为空，最近公共祖先在左子树
	if left != nil {
		return left
	} else { // 右子树查找不为空，最近公共祖先在左子树
		return right
	}
}
