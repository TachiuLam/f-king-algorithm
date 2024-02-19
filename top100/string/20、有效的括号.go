package string

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false


提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

// 利用出栈入栈思想，有括号闭合，就把左括号出栈，否则入栈，最终栈不为0则存在未闭合括号
func isValid(s string) bool {
	var validMap = map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	if len(s)%2 != 0 {
		return false
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if _, ok := validMap[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != validMap[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
