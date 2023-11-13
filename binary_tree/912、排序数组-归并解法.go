package binary_tree

var tmp []int

// 归并排序: Nlog(N)
func sortArray(nums []int) []int {
	tmp = make([]int, len(nums))
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, left, right int) {
	if left == right {
		return
	}
	mid := (left + right) / 2
	sort(nums, left, mid)
	sort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	// 缓存原有nums的值
	for i := left; i <= right; i++ {
		tmp[i] = nums[i]
	}

	// 双指针，合并left和right
	i, j := left, mid+1
	for p := left; p <= right; p++ {
		if i == mid+1 { // left finish
			nums[p] = tmp[j]
			j++
		} else if j == right+1 { // right finish
			nums[p] = tmp[i]
			i++
		} else if tmp[i] > tmp[j] {
			nums[p] = tmp[j]
			j++
		} else {
			nums[p] = tmp[i]
			i++
		}

	}
}
