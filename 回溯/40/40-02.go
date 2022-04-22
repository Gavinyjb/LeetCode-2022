package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int, 0)
	track := make([]int, 0)

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
		for i := start; i < len(candidates); i++ {
			if i > 1 && candidates[i] == candidates[i-1] {
				continue
			}
			//做选择
			track = append(track, candidates[i])
			backtrack(track, target-candidates[i], i+1)
			track = track[:len(track)-1]
		}
	}
	backtrack(track, target, 0)
	return ret
}

func main() {

}
