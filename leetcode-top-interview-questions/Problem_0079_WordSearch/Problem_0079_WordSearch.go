package Problem_0079_WordSearch

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visited := make([][]bool, h)
	for i, _ := range visited {
		visited[i] = make([]bool, w)
	}
	choose := [][]int{
		{0, 1},  //右
		{0, -1}, //左
		{1, 0},  //下
		{-1, 0}, //上
	}
	// 目前到达了b[i][j]，word2bytes[k....]
	// 从b[i][j]出发，能不能搞定word2bytes[k....]  true  false
	var process func(i, j int, k int) bool
	process = func(i, j int, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		visited[i][j] = true
		defer func() { visited[i][j] = false }() // 回溯时还原已访问的单元格
		for _, ints := range choose {
			newI, newJ := i+ints[0], j+ints[1]
			//过往 路径中不存在新的点
			if newI < len(board) && newI >= 0 && newJ < len(board[0]) && newJ >= 0 && !visited[newI][newJ] {
				if process(newI, newJ, k+1) {
					return true
				}
			}
		}

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if process(i, j, 0) {
				return true
			}
		}
	}
	return false
}

////寻找四周目标的坐标
//func findTarget(i, j int, target byte, board [][]byte, path map[string]bool) (int, int, bool) {
//	choose := [][]int{
//		{0, 1},  //右
//		{0, -1}, //左
//		{1, 0},  //下
//		{-1, 0}, //上
//	}
//	for x := 0; x < len(choose); x++ {
//		newI, newJ := i+choose[i][0], j+choose[i][1]
//		key := idxIJ2String(newI, newJ)
//		//过往 路径中不存在新的点
//		if _, ok := path[key]; newI < len(board) && newI >= 0 && newJ < len(board[0]) && newJ >= 0 && !ok {
//			if board[newI][newJ] == target {
//				return newI, newJ, true
//			}
//		}
//	}
//	return -1, -1, false
//}
