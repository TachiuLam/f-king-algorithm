package array

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

// https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/210764/di-k-xiao-shu-jie-fa-ni-zhen-de-dong-ma-by-geek-8m/?envType=featured-list&envId=2cktkvj?envType=featured-list&envId=2cktkvj
// 简化为寻找 第 K 小的元素
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 要找的元素 位置k
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		// 奇书，只需要找一个第k小的点
		return float64(findK(nums1, nums2, length/2+1))
	} else {
		// 偶数长度，需要找第k小和第k-1小求平均值
		return float64(findK(nums1, nums2, length/2)+findK(nums1, nums2, length/2+1)) / 2.0
	}
}

func findK(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0 // 初试两个数组的左右边界都是从0开始
	for {
		if len(nums1) == index1 {
			return nums2[index2+k-1] // nums2的第k个数
		}
		if len(nums2) == index2 {
			return nums1[index1+k-1] // nums1的第k个数
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2]) // 取两数组边界的最小值
		}
		half := k / 2                                 // 从数组中再取一半
		newIndex1 := min(half+index1-1, len(nums1)-1) // nums1右边界
		newIndex2 := min(half+index2-1, len(nums2)-1) // nums2右边界
		if nums1[newIndex1] <= nums2[newIndex2] {     // nums1 新取部分整体小于nums2
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1 // 移动nums1的左边界
		} else {
			k -= newIndex2 - index2 + 1 // 确定新的k
			index2 = newIndex2 + 1      // 移动nums2左边界
		}
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
