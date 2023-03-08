package 数组

// leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
	for num := 0; num < len(s)-1; num++ {
		if num <= (len(s)-1)>>1 {
			s[num], s[len(s)-1-num] = s[len(s)-1-num], s[num]
		} else {
			break
		}
	}
}
