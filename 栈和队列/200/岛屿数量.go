package leetcdoe

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				infect(grid, i, j)
			}
		}
	}
	return res
}

//感染周围为1的地方使其变为2
func infect(grid [][]byte, i, j int) {
	// 如果 i 行 j 列位置已经越界，或者这个位置上不是 1，退出“感染”过程。
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	infect(grid, i+1, j)
	infect(grid, i-1, j)
	infect(grid, i, j+1)
	infect(grid, i, j-1)
}
