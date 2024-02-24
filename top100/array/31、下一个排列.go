package array

/*
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。

示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100
*/
// 双指针
// https://leetcode.cn/problems/next-permutation/solutions/80560/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/?envType=featured-list&envId=2cktkvj?envType=featured-list&envId=2cktkvj
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	length := len(nums)
	i, j, k := length-2, length-1, length-1
	// step1
	for i >= 0 && nums[i] >= nums[j] { // 找到相邻元素，nums[i] < nums[j]的场景
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] { // 从末尾开始，找到大于nums[i]进行替换
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 由step1推导可知，nums[j:]这部分的元素都是降序排列的，需要使其变成生序排列，保障所有元素组成是下一个最小的整数
	for i, j := j, length-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
