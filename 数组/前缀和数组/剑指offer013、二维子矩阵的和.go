package 前缀和数组

type NumMatrix2 struct {
	preSum [][]int
}

func Constructor2(matrix [][]int) NumMatrix2 {
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 {
		return NumMatrix2{}
	}
	// 构造子矩阵和的 二维数组，冗余零列零行
	preSum := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		preSum[i] = make([]int, col)
	}

	// 计算子矩阵和
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix2{
		preSum: preSum,
	}
}

func (this *NumMatrix2) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}
