package leetcode

import (
	"fmt"
	"sort"
)

type arrPlus struct {
	Val    int
	idxOld int
}

func help(arr []int) []*arrPlus {
	ret := make([]*arrPlus, len(arr))
	for idx, v := range arr {
		ret[idx] = &arrPlus{Val: v, idxOld: idx}
	}
	return ret
}
func help1(arrNew []*arrPlus)[][]int{
	ret:=make([][]int,0)
	for i, v := range arrNew {
		temp:=make([]int, 2)
		temp[0]=min(i,v.idxOld)
		temp[1]=max(i,v.idxOld)
		ret = append(ret, temp)
	}
	return ret
}
func help3(duans [][]int) int {
	ret:=0
	sort.Slice(duans,func(i, j int) bool {
		return duans[i][0]<duans[j][0]
	})
	l,r:=-1,-1
	for _, duan := range duans {
		l1,r1:=duan[0],duan[1]
		if l1>=r{
			ret++
			l,r=l1,r1
		}else  {
			r=max(r,r1)
			l=min(l,l1)
		}
	}
	return ret
}
func maxChunksToSorted(arr []int) int {
	arrNew := help(arr)
	sort.SliceStable(arrNew,func(i, j int) bool {
		return arrNew[i].Val<arrNew[j].Val
	})
	for _, v := range arrNew {
		fmt.Println(*v)
	}
	duans:=help1(arrNew)
	res:=help3(duans)
	fmt.Println(res)
	return res

}
func min(a,b int)int  {
	if a<b{
		return a
	}
	return b
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}