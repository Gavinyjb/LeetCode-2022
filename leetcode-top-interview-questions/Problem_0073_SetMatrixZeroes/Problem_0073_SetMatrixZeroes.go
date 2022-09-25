package Problem_0073_SetMatrixZeroes

func setZeroes(matrix [][]int) {
	rowsIsZero := make(map[int]bool, len(matrix))
	colsIsZero := make(map[int]bool, len(matrix[0]))
	rows := len(matrix)
	cols := len(matrix[0])
	zeroIdxX := make(map[int]bool)
	zeroIdxY := make(map[int]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				zeroIdxX[i] = true
				zeroIdxY[j] = true
			}
		}
	}

	for i, _ := range zeroIdxX {
		if !rowsIsZero[i] { //这一行没有感染过
			infectRow(matrix, i, 0)
			rowsIsZero[i] = true
		}
	}
	for j, _ := range zeroIdxY {
		if !colsIsZero[j] { //这一列没有感染过
			infectCol(matrix, 0, j)
			colsIsZero[j] = true
		}
	}
}

//感染一行
func infectRow(matrix [][]int, x, _ int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[x][i] = 0
	}
}

//感染一列
func infectCol(matrix [][]int, _, y int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][y] = 0
	}
}
