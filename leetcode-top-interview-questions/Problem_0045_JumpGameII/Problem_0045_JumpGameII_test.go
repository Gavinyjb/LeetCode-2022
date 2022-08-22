package Problem_0045_JumpGameII

import "testing"

func TestIsJump(t *testing.T) {
	numsList := [][]int{{2, 3, 1, 1, 4}, {2, 1}, {5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}}
	wants := []int{2, 1, 5}
	//numsList:=[][]int{{2,3,1,1,4},{2,1}}
	//wants:=[]int{2,1}
	for i, _ := range wants {
		nums := numsList[i]
		want := wants[i]
		got := jump(nums)
		if want != got {
			t.Errorf("want:%v got:%v\n", want, got)
		}
	}
}
func TestIsJump2(t *testing.T) {
	numsList := [][]int{{2, 3, 1, 1, 4}, {2, 1}, {5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}}
	wants := []int{2, 1, 5}
	//numsList:=[][]int{{2,3,1,1,4},{2,1}}
	//wants:=[]int{2,1}
	for i, _ := range wants {
		nums := numsList[i]
		want := wants[i]
		got := jump2(nums)
		if want != got {
			t.Errorf("want:%v got:%v\n", want, got)
		}
	}
}
