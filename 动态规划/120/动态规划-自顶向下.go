package leetcode

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	// 1、状态定义：f[i][j] 表示从0,0出发，到达i,j的最短路径
	l := len(triangle)
	dp := make([][]int, len(triangle))
	//初始化
	for i := 0; i < l; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		dp[len(triangle)-1][i] = triangle[len(triangle)-1][i]
	}
	//递推求解
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			// 这里分为两种情况：
			// 1、上一层没有左边值
			// 2、上一层没有右边值
			if j-1 < 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j >= len(dp[i-1]) {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
		result := dp[l-1][0]
		for i := 1; i < len(dp[l-1]); i++ {
			result = min(result, dp[l-1][i])
		}
		return result
	}
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
