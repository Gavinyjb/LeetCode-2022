package Problem_0041_FirstMissingPositive

import "testing"

func TestSwap(t *testing.T) {
	nums := []int{7, 8, 9, 11, 12}
	swap(nums, 0, 4)
	want := []int{12, 8, 9, 11, 7}
	for i, _ := range nums {
		if nums[i] != want[i] {
			t.Errorf("want:%v got:%v\n", want, nums)
		}
	}
}
func TestFirst(t *testing.T) {
	//nums:=[]int{7,8,9,11,12}
	nums := []int{2, 1}
	got := firstMissingPositive(nums)
	want := 3
	if got != want {
		t.Errorf("want:%v got:%v\n", want, got)
	}
}
