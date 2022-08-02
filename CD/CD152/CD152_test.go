package CD152

import (
	"fmt"
	"testing"
)

var bestElementsTestdata = []struct {
	in  []int
	k   int
	f   func(nums []int, k int) []int
	out []int
}{
	{[]int{3, 2, 1, 5, 6, 4}, 2, FindBestKElements, []int{6, 5}},
	{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, FindBestKElements, []int{6, 5, 5, 4}},
}

func TestBestElementsLogic(t *testing.T) {
	for _, tt := range bestElementsTestdata {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			out := tt.f(tt.in, tt.k)
			for i, v := range out {
				if v != tt.out[i] {
					t.Errorf("got %q, want %q", out, tt.out)
				}
			}
		})
	}
}
