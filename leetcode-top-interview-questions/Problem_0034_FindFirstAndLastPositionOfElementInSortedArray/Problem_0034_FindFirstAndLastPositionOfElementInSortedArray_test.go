package Problem_0034_FindFirstAndLastPositionOfElementInSortedArray

import "testing"

func TestFirst(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 8, 9, 9, 9, 9, 910}
	target := 7
	got := findFirst(nums, target)
	want := 1
	if got != want {
		t.Errorf("want:%d got:%d\n", want, got)
	}
}
func TestLast(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 8, 9, 9, 9, 9, 910}
	target := 7
	got := findLast(nums, target)
	want := 2
	if got != want {
		t.Errorf("want:%d got:%d\n", want, got)
	}
}
func TestRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 8, 9, 9, 9, 9, 910}
	target := 7
	got := searchRange(nums, target)
	want := []int{1, 2}
	for i := 0; i < 2; i++ {
		if got[i] != want[i] {
			t.Errorf("want:%d got:%d\n", want[i], got[i])
		}
	}

}
