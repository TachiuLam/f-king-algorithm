package binary_tree

import "math/rand"

// FindKthLargest 快速选择 时间复杂度O(N)
func FindKthLargest(nums []int, k int) int {
	index := len(nums) - k // 将k转换为排序后的index
	left, right := 0, len(nums)-1
	for left <= right {
		p := partition(nums, left, right)
		if p == index {
			return nums[index]
		} else if p < index {
			left = p + 1
			for ; left < index && nums[left] == nums[left-1]; left++ {
				//关键代码，跳过与nums[p]相同的值 ,只为了过leetcode有重复测试数据的用例，避免超时
			}
		} else {
			right = p - 1
			for ; right > index && nums[right] == nums[right+1]; right-- {
				//关键代码，跳过与nums[p]相同的值 ,只为了过leetcode有重复测试数据的用例，避免超时
			}
		}
	}
	return -1
}

// 找出排序好的p点的index, 使得nums[left:p-1] <= nums[p] <= nums[p+1, right]
func partition(nums []int, left, right int) int {
	pivotIndex := left + rand.Intn(right-left) // 随机选择一个索引作为基准值
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	p := nums[left]
	i, j := left+1, right
	for i <= j {
		for ; i <= right && nums[i] <= p; i++ { // 最终会找到nums[i] > p 的位置
		}
		for ; j > left && nums[j] > p; j-- { // 最终会找到nums[j] <= p 的位置
		}
		if i > j {
			break
		}
		// swap i，j的值
		nums[i], nums[j] = nums[j], nums[i]
	}
	// j为最终p的index，将它和初始的left进行交换
	nums[left], nums[j] = nums[j], nums[left]
	return j
}
