package Problem_0062_UniquePaths

//组合数计算
//C(n,m)  n:下 m:上  ==n! / (m!*(n-m)!)
func uniquePaths(m int, n int) int {
	part := n - 1
	all := m + n - 2
	//all!/(part!*(all-part)!)

	/*
		all=10
		right=4
		o1= 7 8 9 10
		o2= 1 2 3 4
		o1乘进去的数 一定等于 o2乘进去的个数
	*/
	o1, o2 := 1, 1
	for i, j := part+1, 1; i <= all || j <= all-part; i, j = i+1, j+1 {
		o1 *= i
		o2 *= j
		help := getGreatestCommonDivsor(o1, o2)
		o1 /= help
		o2 /= help
	}
	return o1 / o2
}

// 调用的时候，请保证初次调用时，m和n都不为0
//辗转相除
func gcd(m, n int) int {
	if n == 0 {
		return m
	} else {
		return gcd(n, m%n)
	}
}

//更相减损术  最大公约数
func getGreatestCommonDivsor(m, n int) int {
	big, smaller := max(m, n), min(m, n)
	if smaller == 0 {
		return big
	}
	if big%smaller == 0 {
		return smaller
	}
	return getGreatestCommonDivsor(big-smaller, smaller)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func uniquePaths(m int, n int) int {
//	dp:=make([][]int, m)
//	//初始化
//	for i := 0; i < m; i++ {
//		dp[i]=make([]int, n)
//	}
//	//第一列
//	for i := 0; i < m; i++ {
//		dp[i][0]=1
//	}
//	//第一行
//	for j := 0; j < n; j++ {
//		dp[0][j]=1
//	}
//	for i:=1;i<m;i++{
//		for j := 1; j < n; j++ {
//			dp[i][j]=dp[i-1][j]+dp[i][j-1]
//		}
//	}
//	return dp[m-1][n-1]
//}
