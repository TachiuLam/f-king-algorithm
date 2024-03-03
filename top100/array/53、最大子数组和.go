package array

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
*/
// https://leetcode.cn/problems/maximum-subarray/solutions/847114/dai-ma-sui-xiang-lu-53-zui-da-zi-xu-he-b-xqus/?envType=featured-list&envId=2cktkvj?envType=featured-list&envId=2cktkvj
//时间复杂度：O(n)
//空间复杂度：O(n)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i]+nums[i-1] {
			nums[i] += nums[i-1]
		}
		if maxSum < nums[i] {
			maxSum = nums[i]
		}
	}
	return maxSum
}
