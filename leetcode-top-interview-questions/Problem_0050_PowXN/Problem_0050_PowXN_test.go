package Problem_0050_PowXN

import "testing"

func TestMyPow(t *testing.T) {
	x := 2.00
	n := -2
	want := 0.25
	got := myPow(x, n)
	if want != got {
		t.Errorf("want:%v,got:%v\n", want, got)
	}
}
