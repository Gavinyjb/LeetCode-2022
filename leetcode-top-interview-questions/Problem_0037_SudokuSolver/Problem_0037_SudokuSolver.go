package Problem_0037_SudokuSolver

import "fmt"

func solveSudoku(board [][]byte) {

}

func isValidSudoku(board [][]byte) {
	row := make([]map[byte]bool, 9)
	col := make([]map[byte]bool, 9)
	bucket := make([]map[byte]bool, 9)
	//初始化
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		bucket[i] = make(map[byte]bool)
	}
	initMaps(board, row, col, bucket)
	fmt.Println(row)
	fmt.Println(col)
	fmt.Println(bucket)
	list := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	process(board, 0, 0, row, col, bucket, list)
}
func process(board [][]byte, i, j int, row, col, bucket []map[byte]bool, list []byte) bool {
	if i == 9 {
		return true
	}
	nexti, nextj := -1, -1
	if j != 8 {
		nexti = i
		nextj = j + 1
	} else { //j==8
		nexti = i + 1
		nextj = 0
	}
	if board[i][j] != '.' {
		return process(board, nexti, nextj, row, col, bucket, list)
	} else {
		bucketNum := getBucket(i, j)
		for _, v := range list {
			if !row[i][v] && !col[j][v] && !bucket[bucketNum][v] { //合法才可以
				row[i][v] = true
				col[j][v] = true
				bucket[bucketNum][v] = true
				board[i][j] = v
				if process(board, nexti, nextj, row, col, bucket, list) {
					return true
				}
				row[i][v] = false
				col[j][v] = false
				bucket[bucketNum][v] = false
				board[i][j] = '.'
			}
		}
		return false
	}
}

//保存数独本来的数据
func initMaps(board [][]byte, row, col, bucket []map[byte]bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//fmt.Println(row)
			//fmt.Println(col)
			//fmt.Println(bucket)
			//fmt.Printf("i:%v j:%v\n",i,j)
			bucketNum := getBucket(i, j)
			if board[i][j] != '.' {
				row[i][board[i][j]] = true
				col[j][board[i][j]] = true
				bucket[bucketNum][board[i][j]] = true
			}
		}
	}
}
func getBucket(i, j int) int {
	row, col := 0, 0
	row = i / 3
	col = j / 3
	ret := row*3 + col
	return ret
}
