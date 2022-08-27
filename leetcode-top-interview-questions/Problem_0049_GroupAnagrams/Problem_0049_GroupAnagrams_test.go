package Problem_0049_GroupAnagrams

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	got := groupAnagrams(strs)
	want := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	fmt.Println("want:", want)
	fmt.Println("got:", got)
}
func TestDemo(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	demo(strs)
}
