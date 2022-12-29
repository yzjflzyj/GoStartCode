package main

// NumMatrix
//     * 和剑指 Offer II 013  相同
//     * 304. 二维区域和检索 - 矩阵不可变
//     * 给定一个二维矩阵 matrix，以下类型的多个请求：
//     * 计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1,col1) ，右下角 为 (row2,col2) 。
//     * 实现 NumMatrix 类：
//     * NumMatrix(int[][] matrix)给定整数矩阵 matrix 进行初始化
//     * int sumRegion(int row1, int col1, int row2, int col2)返回 左上角 (row1,col1)、右下角(row2,col2) 所描述的子矩阵的元素 总和 。
type NumMatrix struct {
	preSum [][]int
}

func Constructor1(matrix [][]int) NumMatrix {
	var rows, cols int
	if len(matrix) == 0 {
		rows, cols = 1, 1
	} else {
		rows, cols = len(matrix)+1, len(matrix[0])+1
	}
	preSum := make([][]int, rows)
	for i := range preSum {
		preSum[i] = make([]int, cols)
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{
		preSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}
