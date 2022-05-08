package main

import "sort"

//先分配小的
func findContentChildren(g []int, s []int) int {
	sort.Ints(g) //小朋友
	sort.Ints(s) //饼干
	ret := 0
	for i, j := 0, 0; j < len(s); j++ {
		if i < len(g) {
			if s[j] >= g[i] { // 可以满足
				ret++
				i++
			}
		} else {
			return ret
		}
	}
	return ret
}

func main() {

}
