package string

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/
// hash表+滑动窗口
// 1、利用hash表存储所有元素的出现次数
// 2、当出现次数大于1，则移动滑动窗口左边界，直到满足窗口内元素出现次数小等于1
// 3、更新此时的最大长度
func lengthOfLongestSubstring(s string) int {
	exits := make(map[byte]int)
	left := 0  // 窗口左边界
	right := 0 // 窗口右边界
	maxLength := 0
	for num := range s {
		right++
		exits[s[num]]++
		// 判断窗口左边界是否需要移动
		for exits[s[num]] > 1 {
			d := s[left]
			left++
			exits[d]--
		}
		length := right - left
		if maxLength < length {
			maxLength = length
		}
	}
	return maxLength
}
