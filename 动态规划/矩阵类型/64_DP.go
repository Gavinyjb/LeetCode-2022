package leetcode

func minPathSum(grid [][]int) int {
	dp:=make([][]int,len(grid))

	//初始化
	for i := 0; i < len(dp); i++ {
		dp[i]=make([]int, len(grid[0]))
	}
	dp[len(dp)-1][len(dp[0])-1]=grid[len(grid)-1][len(grid[0])-1]

	for i := len(dp)-2; i >=0; i-- {
		dp[i][len(dp[0])-1]=grid[i][len(dp[0])-1]+dp[i+1][len(dp[0])-1]
	}
	for j := len(dp[0])-2; j >=0 ; j-- {
		dp[len(dp)-1][j]=grid[len(dp)-1][j]+dp[len(dp)-1][j+1]
	}
	for i := len(dp)-2; i >=0; i--{
		for j := len(dp[0])-2; j >=0 ; j--{
			dp[i][j]=grid[i][j]+min(dp[i+1][j],dp[i][j+1])
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