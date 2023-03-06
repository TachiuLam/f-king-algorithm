package 数组

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow + 1
}
