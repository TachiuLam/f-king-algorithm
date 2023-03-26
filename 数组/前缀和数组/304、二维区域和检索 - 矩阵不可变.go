package 前缀和数组

/*
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
*/

type NumMatrix struct {
	// 定义：preSum[i][j] 记录 matrix 中子矩阵 [0, 0, i-1, j-1] 的元素和
	preSum [][]int
}

func ConstructorNum(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 {
		return NumMatrix{}
	}
	// 构造二维数组前缀和，计算0，0，i-1，j-1的前缀和，预留0点，便于计算
	preSum := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		preSum[i] = make([]int, col+1)
	}
	// 计算
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			// 计算每个矩阵 [0, 0, i, j] 的元素和
			preSum[i][j] = preSum[i][j-1] + preSum[i-1][j] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{preSum: preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}
