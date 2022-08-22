package Problem_0042_TrappingRainWater

import (
	"fmt"
	"testing"
)

func TestGreater(t *testing.T) {
	nums := []int{73, 74, 75, 71, 69, 72, 76, 73}
	left := []int{-1, -1, -1, 2, 3, 2, -1, 6}
	right := []int{1, 2, 6, 5, 5, 6, -1, -1}
	gotLeft, gotRight := getNearGreaterNoRepeat(nums)
	fmt.Println(gotLeft, gotRight)
	for i := 0; i < len(nums); i++ {
		if left[i] != gotLeft[i] {
			t.Errorf("want:%v got:%v\n", left, gotLeft)
		}
		if right[i] != gotRight[i] {
			t.Errorf("want:%v got:%v\n", right, gotRight)
		}
	}
}
func TestRrap(t *testing.T) {
	//height:=[]int{0,1,0,2,1,0,1,3,2,1,2,1}
	height := []int{4, 2, 0, 3, 2, 5}
	got := trap(height)
	want := 9
	if want != got {
		t.Errorf("want:%v got:%v\n", want, got)
	}
}

func TestRrap2(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	got := trap2(height)
	want := 6
	if want != got {
		t.Errorf("want:%v got:%v\n", want, got)
	}
}
