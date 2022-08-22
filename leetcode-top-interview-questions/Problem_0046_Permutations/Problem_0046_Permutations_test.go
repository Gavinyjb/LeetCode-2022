package Problem_0046_Permutations

import (
	"fmt"
	"testing"
)

func TestIsJump2(t *testing.T) {
	nums := []int{1, 2, 3}
	wants := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}

	want := wants
	got := permute(nums)
	//if want!=got{
	//	t.Errorf("want:%v got:%v\n",want,got)
	//}
	fmt.Println(want)
	fmt.Println(got)
}
