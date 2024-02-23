package string

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]

提示：

1 <= n <= 8
*/
func generateParenthesis(n int) []string {
	var res = make([]string, 0)
	dfs2(n, 0, 0, "", &res)
	return res
}

func dfs2(n int, lc int, rc int, src string, res *[]string) {
	if lc == n && rc == n {
		*res = append(*res, src)
		return
	}
	if lc < n {
		dfs2(n, lc+1, rc, src+"(", res)
	}
	if lc > rc {
		dfs2(n, lc, rc+1, src+")", res)
	}
}
