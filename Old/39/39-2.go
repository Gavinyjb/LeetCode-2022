package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res:=make([][]int,0)
	track:=make([]int,0)
	sum:=0
	sort.Ints(candidates)
	var backtrack func(track []int,start int,sum int)
	backtrack= func(track []int, start int,sum int) {

		temp:=make([]int,len(track))
		copy(temp,track)
		if sum==target{
			res=append(res,temp )
			return
		}else if sum>target {
			return
		}
		for i := start; i <len(candidates)&&sum+candidates[i]<=target; i++ {

			//做选择
			track = append(track, candidates[i])
			sum+=candidates[i]
			//下一层决策树
			backtrack(track,i,sum)
			track=track[:len(track)-1]// 不用i+1了，表示可以重复读取当前的数
			sum-=candidates[i]
		}
	}
	backtrack(track,0,sum)
	return res

}
func main() {
	a:=[]int{2,3,6,7}
	target :=7
	ret:=combinationSum(a,target)
	fmt.Println(ret)
}
