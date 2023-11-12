package test

import (
	"github.com/TachiuLam/f-king-algorithm/binary_tree"

	"testing"
)

func TestQuickSor(t *testing.T) {
	var nums = []int{6, 1, 2, 5, 3, 4, 0, 0, 2}
	binary_tree.SortArray(nums)
	t.Log(nums)
}

func TestFind(t *testing.T) {
	{
		var nums = []int{6, 1, 2, 5, 3, 4, 0}
		res := binary_tree.FindKthLargest(nums, 1)
		t.Log(res, nums)
	}
	{
		var nums = []int{7, 6, 5, 4, 3, 2, 1}
		res := binary_tree.FindKthLargest(nums, 2)
		t.Log(res, nums)
	}
	{
		var nums = []int{-1, 2, 0}
		res := binary_tree.FindKthLargest(nums, 2)
		t.Log(res, nums)
	}
}
