package 双指针

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	slow, count := 0, 0
	for _, item := range nums {
		if item != 0 {
			nums[slow] = item
			slow++
		} else {
			count++
		}
	}
	for i := 0; i < count; i++ {
		nums[slow+i] = 0
	}
}
