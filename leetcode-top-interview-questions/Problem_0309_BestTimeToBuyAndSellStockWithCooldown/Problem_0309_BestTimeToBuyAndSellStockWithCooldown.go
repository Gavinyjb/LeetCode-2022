package Problem_0309_BestTimeToBuyAndSellStockWithCooldown

func maxProfit(prices []int) int {
	return 0
}

// buy == false 目前可以交易，而且当前没有购买行为
// buy == true 已经买了，买入价buyPrices，待卖出
func try(prices []int, buy bool, index int, buyPrice int) int {
	if index >= len(prices) {
		return 0
	}
	if buy {
		noSell := try(prices, true, index+1, buyPrice)
		yesSell := prices[index] - buyPrice + try(prices, false, index+2, 0)
		return max(noSell, yesSell)
	} else {
		noBuy := try(prices, false, index+1, 0)
		yesBuy := try(prices, true, index+1, prices[index])
		return max(noBuy, yesBuy)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
