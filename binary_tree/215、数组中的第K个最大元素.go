package binary_tree

func FindKthLargest(nums []int, k int) int {
	index := len(nums) - k // 将k转换为排序后的index
	quickSortByK(nums, index)
	return nums[index]
}

func quickSortByK(nums []int, index int) {
	if len(nums) <= 1 {
		return
	}
	p := nums[0]
	i, j := -1, len(nums)
	for i < j {
		for i++; nums[i] < p; i++ {
		}
		for j--; nums[j] > p; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if index >= j+1 {
		index -= j + 1 //index对于nums[j+1:]，位置已经改变
		quickSortByK(nums[j+1:], index)
	} else {
		quickSortByK(nums[:j+1], index)
	}
	//quickSortByK(nums[j+1:], k)
	//quickSortByK(nums[:j+1], k)
}
