package Problem_0033_SearchInRotatedSortedArray

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	got := search(nums, target)
	want := 4
	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}
