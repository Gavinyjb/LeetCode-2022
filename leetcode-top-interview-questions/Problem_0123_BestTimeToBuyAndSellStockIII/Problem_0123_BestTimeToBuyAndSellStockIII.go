package Problem_0123_BestTimeToBuyAndSellStockIII

//输入：prices = [3,3,5,0,0,3,1,4]
//输出：6
//解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3
//
//链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii
func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	ans := 0
	doneOnceMinusBuyMax := -prices[0]
	doneOnceMax := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		ans = maxFunc(ans, doneOnceMinusBuyMax+prices[i])
		min = minFunc(min, prices[i])
		doneOnceMax = maxFunc(doneOnceMax, prices[i]-min)
		doneOnceMinusBuyMax = maxFunc(doneOnceMinusBuyMax, doneOnceMax-prices[i])
	}
	return ans
}
func maxFunc(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func minFunc(a, b int) int {
	if a > b {
		return b
	}
	return a
}
