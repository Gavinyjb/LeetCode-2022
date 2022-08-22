package Problem_0029_DivideTwoIntegers

import "testing"

func TestAdd(t *testing.T) {
	sum := add(1, 2)
	wants := 3
	if sum != wants {
		t.Errorf("expected '%d' but got '%d'", wants, sum)
	}
}
func TestNegnum(t *testing.T) {
	out := negNum(3)
	wants := -3
	if out != wants {
		t.Errorf("expected '%d' but got '%d'", wants, out)
	}
}

//减法测试
func TestMinus(t *testing.T) {
	out := minus(2, 4)
	wants := -2
	if out != wants {
		t.Errorf("expected '%d' but got '%d'", wants, out)
	}
}

func TestMulti(t *testing.T) {
	got := multi(2, 4)
	want := 8
	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func TestDiv(t *testing.T) {
	got1, got2 := div(9, 2)
	want1, want2 := 4, 1
	if got1 != want1 {
		t.Errorf("expected '%d' but got '%d'", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("expected '%d' but got '%d'", want2, got2)
	}
}
