package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(candidates))
	sum := 0
	sort.Ints(candidates)
	var backtrack func(track []int, start int, sum int, used []bool)
	backtrack = func(track []int, start int, sum int, used []bool) {

		temp := make([]int, len(track))
		copy(temp, track)
		if sum == target {
			res = append(res, temp)
			return
		} else if sum > target {
			return
		}
		for i := start; i < len(candidates) && sum+candidates[i] <= target; i++ {
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				continue
			}
			//做选择
			track = append(track, candidates[i])
			used[i] = true
			sum += candidates[i]
			//下一层决策树
			backtrack(track, i+1, sum, used)
			track = track[:len(track)-1] // 不用i+1了，表示可以重复读取当前的数
			used[i] = false
			sum -= candidates[i]
		}
	}
	backtrack(track, 0, sum, used)
	return res
}
func main() {
	a := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	ret := combinationSum2(a, target)
	fmt.Println(ret)

}
