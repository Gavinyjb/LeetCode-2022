package Problem_0188_BestTimeToBuyAndSellStockIV

//如果k>len(prices)/2  退化成可以无限次交易，
//因为一个数组的上坡阶段最多有len(prices)/2次，我们最多也就是需要做这么多次交易
func maxProfit1(k int, prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	if k >= len(prices)/2 {
		return allProfit(prices)
	}
	dp := make([][]int, len(prices))
	//初始化
	for i, _ := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			//1)  第i天不发生第j次交易
			dp[i][j] = dp[i-1][j]
			//2)  第i天发生第j次交易，那么这次交易卖出的股票是第p天买入的 也就意味着前j-1次交易发生在p天前
			for p := 0; p <= i; p++ {
				dp[i][j] = max(dp[p][j-1]+prices[i]-prices[p], dp[i][j])
			}
		}
	}
	return dp[len(prices)-1][k]
}

//斜率优化后
func maxProfit2(k int, prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	if k >= len(prices)/2 {
		return allProfit(prices)
	}
	dp := make([][]int, len(prices))
	//初始化
	for i, _ := range dp {
		dp[i] = make([]int, k+1)
	}
	ans := 0
	for j := 1; j <= k; j++ {
		pre := dp[0][j]
		best := pre - prices[0]
		for i := 1; i < len(prices); i++ {
			pre = dp[i][j-1]
			dp[i][j] = max(dp[i-1][j], prices[i]+best)
			best = max(best, pre-prices[i])
			ans = max(dp[i][j], ans)
		}
	}
	return ans
}

func allProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ret := 0
	for i := 1; i < len(prices); i++ {
		ret += max(0, prices[i]-prices[i-1])
	}
	return ret
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
