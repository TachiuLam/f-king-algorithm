package 数组

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
限制：

1 <= target <= 10^5
*/
func findContinuousSequence(target int) [][]int {
	var res [][]int
	for i := 1; i < target; i++ {
		child := findChildSequence(target, i)
		if len(child) > 1 {
			res = append(res, child)
		}
	}
	return res

}

func findChildSequence(target, n int) []int {
	var (
		res []int
		sum int
	)
	for i := n; i < target; i++ {
		sum += i
		res = append(res, i)
		if sum == target {
			return res
		} else if sum > target {
			return []int{}
		}
	}
	return []int{}
}

// 双指针优化
func findContinuousSequence2(target int) [][]int {
	var res [][]int
	for l, r := 1, 2; l < r; {
		sum := (l + r) * (r - l + 1) / 2 // 区间求和
		if sum == target {
			var c []int
			for i := l; i <= r; i++ {
				c = append(c, i)
			}
			res = append(res, c)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return res

}
