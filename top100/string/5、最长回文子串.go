package string

/*
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/
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

func findPalindrome(s string, left, right int) string {
	if len(s) <= 1 {
		return s
	}
	// 由中间向两边，查找最长回文串
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 左闭右开，需要将左右都减去一个元素
	return s[left+1 : right]
}
