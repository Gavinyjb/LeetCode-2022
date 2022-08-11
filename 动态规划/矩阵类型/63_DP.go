package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if le
	rows:=len(obstacleGrid)
	cols:=len(obstacleGrid[0])
	dp:=make([][]int,len(obstacleGrid))
	//初始化
	for i, _ := range dp {
		dp[i]=make([]int, len(obstacleGrid[0]))
	}
	// 原点
	if obstacleGrid[0][0]==0{
		dp[0][0]=1
	}else {
		dp[0][0]=0
	}
	//行
	for i := 1; i < cols; i++ {
		if obstacleGrid[0][i]==1||dp[0][i-1]==0{
			dp[0][i]=0
		}else {
			dp[0][i]=1
		}
	}
	//列
	for i := 1; i < rows;i++ {
		if obstacleGrid[i][0]==1||dp[i-1][0]==0{
			dp[i][0]=0
		}else {
			dp[i][0]=1
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if obstacleGrid[i][j]==1{
				dp[i][j]=0
			}else {
				dp[i][j]=dp[i-1][j]+dp[i][j-1]
			}
		}
	}
	return dp[rows-1][cols-1]
}