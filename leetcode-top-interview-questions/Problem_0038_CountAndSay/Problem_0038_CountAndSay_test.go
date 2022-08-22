package Problem_0038_CountAndSay

import "testing"

func TestNtoSay(t *testing.T) {
	n := 4
	got := countAndSay(n)
	want := "1211"
	if got != want {
		t.Errorf("want:%s got:%s\n", want, got)
	}
}
