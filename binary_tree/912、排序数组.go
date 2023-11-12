package binary_tree

// SortArray 快速排序: 先排序好且分点位置p在数组中的位置，再找下一个p，排序该p在数组的位置
func SortArray(nums []int) []int {
	quickSort(nums)
	return nums
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	p := nums[0] // 也可以使用rand包随机选择
	left, right := -1, len(nums)
	for left < right {
		// 找到 p >= nums[left] 的位置，for循环开始时left++，所以初始的left为-1
		for left++; nums[left] < p; left++ {
		}
		for right--; nums[right] > p; right-- {
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	quickSort(nums[:right+1])
	quickSort(nums[right+1:])
}
