package binary_tree

var indexValMap map[int]int

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	indexValMap = make(map[int]int, len(preorder))
	for index, val := range postorder {
		indexValMap[val] = index
	}
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := build(preorder, postorder, 0, len(preorder)-1, 0, len(postorder)-1)
	return root
}

// preLeft，preRight分别是前序数组的左右边界，postLeft，postRight是后序数组的左右边界
func build(preorder []int, postorder []int, preLeft, preRight, postLeft, postRight int) *TreeNode {
	// 构造左子树
	if preLeft > preRight {
		return nil
	}
	root := &TreeNode{Val: preorder[preLeft]}
	if preLeft == preRight {
		return root
	}
	// preorder[preLeft+1]是左子树的值，找到该值在后序数组中的索引
	leftRootPostIndex := indexValMap[preorder[preLeft+1]]
	leftLength := leftRootPostIndex + 1 - postLeft
	root.Left = build(preorder, postorder, preLeft+1, leftLength+preLeft, postLeft, postLeft+leftLength-1)
	root.Right = build(preorder, postorder, preLeft+leftLength+1, preRight, postLeft+leftLength, postRight-1)
	return root
}
