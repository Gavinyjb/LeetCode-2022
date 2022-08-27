package Problem_0123_BestTimeToBuyAndSellStockIII

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}

	want := 6
	got := maxProfit(prices)
	if got != want {
		t.Errorf("want:=%v got:=%v\n", want, got)
	}
}
