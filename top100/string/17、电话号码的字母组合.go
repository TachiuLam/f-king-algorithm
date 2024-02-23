package string

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]

提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/?envType=featured-list&envId=2cktkvj?envType=featured-list&envId=2cktkvj

func letterCombinations(digits string) []string {
	var (
		m         = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"} // 2-9
		path, res = make([]byte, 0), make([]string, 0)
	)
	if digits == "" {
		return res
	}
	dfs(m, &res, &path, digits, 0)
	return res
}

func dfs(m []string, res *[]string, path *[]byte, digits string, start int) {
	if len(*path) == len(digits) {
		tmp := string(*path)
		*res = append(*res, tmp)
		return
	}
	digit := int(digits[start] - '0')
	str := m[digit-2]
	for i := 0; i < len(str); i++ {
		*path = append(*path, str[i])
		dfs(m, res, path, digits, start+1)
		*path = (*path)[:len(*path)-1] // 恢复此次append的元素，为下个for循环使用
	}
}
