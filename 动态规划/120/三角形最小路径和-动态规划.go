package leetcode

//动态规划 -自底向上
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	//状态定义 dp[i][j]表示从i，j出发，到达最后一层所需要的最短路径
	l := len(triangle)
	dp := make([][]int, l)
	//初始化
	for i := 0; i < l; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		dp[len(triangle)-1][i] = triangle[len(triangle)-1][i]
	}
	//递推求解
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
