package array

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/
// 前后指针法
// 前指针右移，后指针左移
// 1、如果前缀最大值比后缀最大值的高度低，那么左指针位置的右侧模板高度则一定会高于左侧木板，因此一定能接到雨水，该位置接的雨水就是较矮的木板（最大左前缀）和left本身高度的差
// 2、右侧同理
// 3、直到左右指针相遇，遍历结束，相遇点可不判断，因为最大值一定不能接雨水
//复杂度分析
//时间复杂度：O(n)\mathcal{O}(n)O(n)，其中 nnn 为 height\textit{height}height 的长度。
//空间复杂度：O(1)\mathcal{O}(1)O(1)，仅用到若干变量。
func trap(height []int) int {
	left := 0
	right := len(height) - 1
	maxPre, maxEnd := 0, 0 // 前缀 后缀最大值初始化都是0
	res := 0
	for left < right {
		maxPre = max(maxPre, height[left])  // 计算前缀最大值
		maxEnd = max(maxEnd, height[right]) // 计算后缀最大值

		if maxPre < maxEnd {
			res += maxPre - height[left]
			left++
		} else {
			res += maxEnd - height[right]
			right--
		}
	}
	return res
}
