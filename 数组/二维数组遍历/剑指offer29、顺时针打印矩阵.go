package 二维数组遍历

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		index                    = 0
		res                      = make([]int, rows*columns)
		left, right, bottom, top = 0, columns - 1, rows - 1, 0
	)
	for left <= right && top <= bottom {
		// 从左到右
		for i := left; i <= right; i++ {
			res[index] = matrix[top][i]
			index++
		}
		// 从上到下
		for i := top + 1; i <= bottom; i++ {
			res[index] = matrix[i][right]
			index++
		}
		if left < right && top < bottom {
			// 从右到左
			for i := right - 1; i > left; i-- {
				res[index] = matrix[bottom][i]
				index++
			}
			// 从下到上
			for i := bottom; i > top; i-- {
				res[index] = matrix[i][left]
				index++
			}
		}
		left++
		top++
		right--
		bottom--
	}
	return res
}
