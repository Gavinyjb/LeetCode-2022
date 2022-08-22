package Problem_0144_BinaryTreePreOrderTraversal

import "testing"

func TestMorris(t *testing.T) {
	rootInts := []int{1, -1 << 63, 2, 3}
	root := Ints2TreeNode(rootInts)
	want := []int{1, 2, 3}
	got := morrisPre(root)
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
	want := []int{1, 2, 3}
	got := preorderTraversal(root)
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
