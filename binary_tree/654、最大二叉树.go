package binary_tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var maxIndex int
	for index, each := range nums {
		if each > nums[maxIndex] {
			maxIndex = index
		}
	}
	// 构造根节点
	root := &TreeNode{Val: nums[maxIndex]}
	// 构造左子树
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	// 构造右子树
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}
