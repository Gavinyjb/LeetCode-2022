package Problem_0036_ValidSudoku

func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]bool, 9)
	col := make([]map[byte]bool, 9)
	bucket := make([]map[byte]bool, 9)
	//初始化
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		bucket[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//fmt.Println(row)
			//fmt.Println(col)
			//fmt.Println(bucket)
			//fmt.Printf("i:%v j:%v\n",i,j)
			bucketNum := getBucket(i, j)
			if board[i][j] != '.' {
				if _, ok := row[i][board[i][j]]; ok { //行存在
					return false
				} else {
					row[i][board[i][j]] = true
				}
				if _, ok := col[j][board[i][j]]; ok { //列存在
					return false
				} else {
					col[j][board[i][j]] = true
				}
				if _, ok := bucket[bucketNum][board[i][j]]; ok { //桶存在
					return false
				} else {
					bucket[bucketNum][board[i][j]] = true
				}
			}
		}
	}
	return true
}
func getBucket(i, j int) int {
	row, col := 0, 0
	row = i / 3
	col = j / 3
	ret := row*3 + col
	return ret
}
