package string

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。

示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

提示：

m == s.length
n == t.length
1 <= m, n <= 105
s 和 t 由英文字母组成
*/
//解题思路
//思路：滑动窗口+哈希表(计数器)
//
//先遍历 t， 记录元素到哈希表 need 中；记录需要遍历的个数 len(t) 到 neecCnt。
//初始化长度为2的 ret 数组[0, Inf]，作为记录目标子字符串的左右索引下标。
//遍历字符串 s， 如果遍历到的元素刚好>0, 说明是t中元素，needCnt减一。将遍历到的元素-1记录到哈希表 need 中。
//如果 needCnt == 0 说明此时左右下标范围内 [lo,hi] 有符合要求子字符串, 开始缩小滑动窗口。
//左索引下标 lo 向前推进, 当 need[s[lo]] == 0 时，说明遍历到t中元素（并且因为此时 needCntCnt==0）， s[lo,hi]为结果之一，判断是否最小。
// https://leetcode.cn/problems/minimum-window-substring/submissions/509450298/?envType=featured-list&envId=2cktkvj?envType=featured-list&envId=2cktkvj
func minWindow(s string, t string) string {
	need := make(map[byte]int) // 记录遍历的元素次数
	for i := range t {
		need[t[i]]++ // 记录各个目标元素需要出现的次数
	}
	needCnt := len(t)       // 目标字符总数
	ret := [2]int{0, 16384} // s子串的左右边界
	var lo = 0
	for hi := range s {
		if need[s[hi]] > 0 {
			needCnt--
		}
		need[s[hi]]--     // 遍历到的元素都需要记录-1次
		if needCnt == 0 { // 子串匹配了所有t的字符，需要判断是否可以左移滑动窗口的左边界，寻找更小的子串
			for true {
				if need[s[lo]] == 0 { // 说明是需要的目标元素刚好达到需要的次数，因此不能左移窗口
					break
				}
				need[s[lo]]++ // 移出的元素需要+1，用于判断是否移除了需要的元素
				lo++
			}
			// 判断是否最小子串
			if hi-lo < ret[1]-ret[0] {
				ret = [2]int{lo, hi}
			}
			// 继续尝试左移滑动窗口，查找是否有更小子串
			need[s[lo]]++
			lo++
			needCnt++
		}
	}
	if ret[1] > len(s) {
		return ""
	}
	return s[ret[0] : ret[1]+1]

}
