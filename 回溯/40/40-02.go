package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(candidates))

	var backtrack func(track []int, target int, start int)
	backtrack = func(track []int, target int, start int) {
		if target == 0 {
			temp := make([]int, len(track))
			copy(temp, track)
			ret = append(ret, temp)
		}
		if target < 0 {
			return
		}
		fmt.Println("track:", track)
		fmt.Println("used:", used)
		for i := start; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				continue
			}
			//做选择
			track = append(track, candidates[i])
			used[i] = true
			backtrack(track, target-candidates[i], i+1)
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrack(track, target, 0)
	return ret
}
func main() {
	input := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	ret := combinationSum2(input, target)
	fmt.Println(ret)
}
