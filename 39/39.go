package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res:=make([][]int,0)
	track:=make([]int,0)
	var backtrack func(track []int,start int)
	backtrack= func(track []int, start int) {
		sum:=0
		temp:=make([]int,len(track))
		copy(temp,track)
		for _, v := range temp {
			sum+=v
		}
		if sum==target{
			res=append(res,temp )
			return
		}else if sum>target {
			return
		}
		for i := start; i <len(candidates) ; i++ {
			//做选择
			track = append(track, candidates[i])
			//下一层决策树
			backtrack(track,i)
			track=track[:len(track)-1]// 不用i+1了，表示可以重复读取当前的数
		}
	}
	backtrack(track,0)
	return res

}
func main() {
	a:=[]int{2,3,6,7}
	target :=7
	ret:=combinationSum(a,target)
	fmt.Println(ret)
}
