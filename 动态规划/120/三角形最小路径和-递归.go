package leetcode

func minimumTotal(triangle [][]int) int {
	var best int = 1 << 31
	dfs(0, 0, 0, triangle, &best)
	return best
}

func dfs(x, y, sum int, a [][]int, best *int) {
	if x == len(a) {
		if sum < *best {
			*best = sum
		}
		return
	}
	dfs(x+1, y, sum+a[x][y], a, best)
	dfs(x+1, y+1, sum+a[x][y], a, best)
}
