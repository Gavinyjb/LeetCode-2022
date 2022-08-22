package Problem_0121_BestTimeToBuyAndSellStock

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}

	want := 5
	got := maxProfit(prices)
	if got != want {
		t.Errorf("want:=%v got:=%v\n", want, got)
	}
}
