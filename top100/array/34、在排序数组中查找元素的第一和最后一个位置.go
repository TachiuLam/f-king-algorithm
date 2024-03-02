package array

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/
//进行两次边界查找
//左边界查找确定左节点
//右边界查找确定右节点
// 时间复杂度 O(log n) 空间复杂度O(1)
func searchRange(nums []int, target int) []int {
	// find left
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	num1 := left
	if left < 0 || left > len(nums)-1 || nums[left] != target {
		num1 = -1
	}

	// find right
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target >= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	num2 := right
	if right < 0 || right > len(nums)-1 || nums[right] != target {
		num2 = -1
	}
	return []int{num1, num2}
}
