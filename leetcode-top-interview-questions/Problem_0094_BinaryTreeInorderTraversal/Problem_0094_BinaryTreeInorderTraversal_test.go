package Problem_0094_BinaryTreeInorderTraversal

import "testing"

func TestMorris(t *testing.T) {
	rootInts := []int{1, -1 << 63, 2, 3}
	root := Ints2TreeNode(rootInts)
	want := []int{1, 3, 2}
	got := morrisIn(root)
	if len(want) != len(got) {
		t.Errorf("want:%v got:%v\n", want, got)
	} else {
		for i, _ := range got {
			if got[i] != want[i] {
				t.Errorf("want:%v got:%v\n", want, got)
			}
		}
	}
}
func TestName(t *testing.T) {
	rootInts := []int{1, -1 << 63, 2, 3}
	root := Ints2TreeNode(rootInts)
	want := []int{1, 3, 2}
	got := inorderTraversal(root)
	if len(want) != len(got) {
		t.Errorf("want:%v got:%v\n", want, got)
	} else {
		for i, _ := range got {
			if got[i] != want[i] {
				t.Errorf("want:%v got:%v\n", want, got)
			}
		}
	}
}
