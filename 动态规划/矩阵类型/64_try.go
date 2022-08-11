package leetcode

import "math"

func minPathSum_1(grid [][]int) int {
	return help(grid, 0, 0)
}

//坐标x,y 到达右下角的最小路径和
func help_1(grid [][]int, x, y int) int {
	if x==len(grid)-1&&y==len(grid[0])-1{
		return grid[x][y]
	}
	if x == len(grid)-1 && y >= 0 && y < len(grid[0])-1 {
		return grid[x][y] + help(grid, x, y+1)
	}
	if y == len(grid[0])-1 && x >= 0 && x < len(grid)-1 {
		return grid[x][y] + help(grid, x+1, y)
	}
	// if y >= 0 && y < len(grid[0]) && x >= 0 && x < len(grid) {
	// 	return grid[x][y] + min(help(grid, x+1, y), help(grid, x, y+1))
	// }
	return grid[x][y] + min(help(grid, x+1, y), help(grid, x, y+1))
}
func min_1(a, b int) int {
	if a < b {
		return a
	}
	return b
}