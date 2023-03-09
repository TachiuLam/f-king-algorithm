package 数组

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)
	for left < right {
		if numbers[left]+numbers[right-1] == target {
			return []int{left + 1, right}
		} else if numbers[left]+numbers[right-1] < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
