package Problem_0044_WildcardMatching

import "testing"

func TestIsMatch(t *testing.T) {
	sList := []string{"cb", "aa", "aa", "abbabbbaabaaabbbbbabbabbabbbabbaaabbbababbabaaabbab"}
	pList := []string{"?a", "a", "*", "*aabb***aa**a******aa*"}
	wants := []bool{false, false, true, true}
	//sList:=[]string{"aa"}
	//pList:=[]string{"*"}
	//wants:=[]bool{true}
	for i, _ := range wants {
		s := sList[i]
		p := pList[i]
		want := wants[i]
		got := isMatch(s, p)
		if want != got {
			t.Errorf("want:%v got:%v\n", want, got)
		}
	}

}
func TestIsMatch2(t *testing.T) {
	sList := []string{"cb", "aa", "aa", "abbabbbaabaaabbbbbabbabbabbbabbaaabbbababbabaaabbab"}
	pList := []string{"?a", "a", "*", "*aabb***aa**a******aa*"}
	wants := []bool{false, false, true, true}
	//sList:=[]string{"aa"}
	//pList:=[]string{"*"}
	//wants:=[]bool{true}
	for i, _ := range wants {
		s := sList[i]
		p := pList[i]
		want := wants[i]
		got := isMatch2(s, p)
		if want != got {
			t.Errorf("want:%v got:%v\n", want, got)
		}
	}
}
func BenchmarkMatch(b *testing.B) {
	//cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
	//BenchmarkMatch
	//BenchmarkMatch-8          190018              5613 ns/op
	for i := 0; i < b.N; i++ {
		sList := []string{"cb", "aa", "aa", "abbabbbaabaaabbbbbabbabbabbbabbaaabbbababbabaaabbab"}
		pList := []string{"?a", "a", "*", "*aabb***aa**a******aa*"}
		wants := []bool{false, false, true, true}
		//sList:=[]string{"aa"}
		//pList:=[]string{"*"}
		//wants:=[]bool{true}
		for i, _ := range wants {
			s := sList[i]
			p := pList[i]
			want := wants[i]
			got := isMatch(s, p)
			if want != got {
				b.Errorf("want:%v got:%v\n", want, got)
			}
		}
	}
}
func BenchmarkMatch2(b *testing.B) {
	//cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
	//BenchmarkMatch2
	//BenchmarkMatch2-8              1        6967060500 ns/op
	for i := 0; i < b.N; i++ {
		sList := []string{"cb", "aa", "aa", "abbabbbaabaaabbbbbabbabbabbbabbaaabbbababbabaaabbab"}
		pList := []string{"?a", "a", "*", "*aabb***aa**a******aa*"}
		wants := []bool{false, false, true, true}
		//sList:=[]string{"aa"}
		//pList:=[]string{"*"}
		//wants:=[]bool{true}
		for i, _ := range wants {
			s := sList[i]
			p := pList[i]
			want := wants[i]
			got := isMatch2(s, p)
			if want != got {
				b.Errorf("want:%v got:%v\n", want, got)
			}
		}
	}
}
