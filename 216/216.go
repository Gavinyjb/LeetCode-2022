package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	nums := make([]int, 9)
	for i := 0; i < 9; i++ {
		nums[i] = i + 1
	}
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(track []int, start int)
	backtrack = func(track []int, start int) {
		//终止条件
		if len(track) == k {
			temp := make([]int, k)
			copy(temp, track)
			sum := 0
			for _, v := range temp {
				sum += v
			}
			if sum == n {
				res = append(res, temp)
			}
		}
		if len(track)+n-start+1 < k {
			return
		}
		//for 选择列表 组合问题，选过了的不再重复选
		for i := start; i < 9; i++ {
			//做选择
			track = append(track, nums[i])
			//进入下一层决策树
			backtrack(track, i+1)
			track = track[:len(track)-1]
		}
	}
	backtrack(track, 0)
	return res
}
func main() {
	fmt.Println(combinationSum3(3,9))

}
