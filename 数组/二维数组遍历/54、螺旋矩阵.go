package 二维数组遍历

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/

func spiralOrder2(matrix [][]int) []int {
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
