package Problem_0070_ClimbingStairs

func climbStairs(n int) int {
	//return help(n)
	return dp(n)
}

/*
1）可以优化空间使用有限几个变量
2）可以考虑矩阵运算加速
*/

func dp(n int) int {
	dp := make([]int, n+1)
	if n < 2 {
		return 1
	}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//一次可以跳1,2。。。m次
//在idx位置到达n 有几种方法
func try(n, m int, idx int) int {
	if idx == n {
		return 1 //只有一种 不跳
	}
	if idx+1 == n {
		return 1 //只有一种跳一次
	}
	sum := 0
	//可以选择 跳一次或者跳。。。m次
	for i := 1; i <= m && idx+i <= n; i++ {
		sum += try(n, m, idx+i)
	}
	return sum
}
func dpPlus(n, m int) int {
	dp := make([]int, n+1)
	if n < 2 {
		return 1
	}
	dp[n] = 1
	dp[n-1] = 1
	for idx := n - 2; idx >= 0; idx-- {
		sum := 0
		//可以选择 跳一次或者跳。。。m次
		for i := 1; i <= m && idx+i <= n; i++ {
			sum += dp[idx+i]
		}
		dp[idx] = sum
	}
	return dp[0]
}

//剩余n 级台阶
func help(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return help(n-1) + help(n-2)
}
