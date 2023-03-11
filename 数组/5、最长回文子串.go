package 数组

// leetcode submit region begin(Prohibit modification and deletion)
// for 0 <= i < len(s):
//
//	找到以 s[i] 为中心的回文串
//	找到以 s[i] 和 s[i+1] 为中心的回文串
//	更新最长的回文字串答案
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s)-1; i++ {
		// 奇数回文子串
		s1 := findPalindrome(s, i, i)
		// 偶数回文子串
		s2 := findPalindrome(s, i, i+1)

		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

// 找到字符串中的回文字串(左右指针 由中间向两端延伸)
func findPalindrome(s string, left, right int) string {
	for (left >= 0) && (right < len(s)) && s[left] == s[right] {
		left--
		right++
	}
	// 左闭右开
	return s[left+1 : right]
}
