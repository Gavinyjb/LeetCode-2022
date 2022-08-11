package leetcode

import "math"

func minimumTotal_1(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	return help(triangle, 0, 0)
}

func help_1(triangle [][]int, x, y int) int {
	if y>=0&&y<len(triangle[len(triangle)-1])&& x==len(triangle)-1{
		return triangle[x][y]
	}
	//越界
	if x >= len(triangle) || x < 0 || y < 0 || y >= len(triangle[x]) {
		return math.MaxInt
	}
	return triangle[x][y] + min(help(triangle, x+1, y), help(triangle, x+1, y+1))
}
func min_1(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}