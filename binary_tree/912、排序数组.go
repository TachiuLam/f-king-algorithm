package binary_tree

import "math/rand"

// SortArray 快速排序: 先排序好且分点位置p在数组中的位置，再找下一个p，排序该p在数组的位置
func SortArray(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	left, right := 0, length-1
	quickSort(nums, left, right)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	p := partitionP(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

func partitionP(nums []int, left, right int) int {
	pivotIndex := left + rand.Intn(right-left) // 随机选择一个索引作为基准值
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	p := nums[left]

	i, j := left+1, right
	for i <= j {
		for ; i <= right && nums[i] <= p; i++ {
		}
		for ; j > left && nums[j] > p; j-- {
		}
		if i > j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}
