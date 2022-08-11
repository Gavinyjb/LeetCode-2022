package leetcode

func minimumTotal(triangle [][]int) int {
	if len(triangle)==0{
		return 0
	}
	dp:=make([][]int,len(triangle),len(triangle))
	//初始化
	for i:=0;i<len(triangle);i++{
		dp[i]=make([]int, len(triangle[i]))
	}
	
	for i := 0; i <len(triangle[len(triangle)-1]); i++ {
		dp[len(triangle)-1][i]=triangle[len(triangle)-1][i]
	}
	for i := len(triangle)-2; i >=0; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j]=triangle[i][j]+min(dp[i+1][j],dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}