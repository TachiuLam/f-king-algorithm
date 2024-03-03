package array

/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

示例 2：
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

提示：

n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000

*/
// 解题思路
//顺时针旋转90：先沿对角线反转矩阵，再沿竖中轴线反转矩阵；
//顺时针旋转180：先沿横中轴线反转矩阵，再沿竖中轴线反转矩阵；
//顺时针旋转270：先沿对角线反转矩阵，再沿横中轴线反转矩阵；
func rotate(matrix [][]int) {
	idx, halfLen := len(matrix)-1, len(matrix)/2 // 记录最后一个元素的横坐标和边长的一半
	// 按对角线反转矩阵 (互换横纵坐标即可)
	for i := range matrix {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 按竖中轴线反转矩阵（交换横坐标，纵坐标不变）
	for i := range matrix {
		for j := 0; j < halfLen; j++ {
			matrix[i][j], matrix[i][idx-j] = matrix[i][idx-j], matrix[i][j]
		}
	}

	// 按横中轴线反转矩阵（交换纵坐标，横坐标不变）
	//for i := 0; i < halfLen; i++ {
	//	for j := range matrix[i] {
	//		matrix[idx-i][j], matrix[i][j] = matrix[i][j], matrix[idx-i][j]
	//	}
	//}
}
