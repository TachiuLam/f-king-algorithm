package 二维数组遍历

/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
*/

// 沿对角线翻转矩阵，再反转每一行
func rotate(matrix [][]int) {
	n := len(matrix)
	// 对角线翻转矩阵
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	// 反转每一行
	for _, row := range matrix {
		reverse(row)
	}
}

func reverse(m []int) {
	i, j := 0, len(m)-1
	for i < j {
		temp := m[i]
		m[i] = m[j-i]
		m[j-i] = temp
		i++
		j--
	}
}
