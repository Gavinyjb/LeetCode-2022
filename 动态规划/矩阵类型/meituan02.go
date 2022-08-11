package main

import (
	"fmt"
	"math"
)

// func help(upNums,downNums map[int]int,)
func main() {
	n := 0
	fmt.Scanf("%v\n",&n)
	mid:=help(n)
	// up:=make(map[int]int)
	// down:=make(map[int]int)
	upList:=make([]int,n)
	downLsit:=make([]int,n)

	for i := 0; i < n; i++ {
		fmt.Scan(&upList[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&downLsit[i])
	}

	upValNums:=make(map[int]int)
	downValNums:=make(map[int]int)
	same:=make(map[int]int)
	for i := 0; i < n; i++ {
		if upList[i]==downLsit[i]{
			same[upList[i]]++
		}else{
			upValNums[upList[i]]++
			downValNums[downLsit[i]]++
		}
	}
	
	//魔法种类集合
	magicSet:=make(map[int]bool)
	for i := 0; i < n; i++ {
		magicSet[upList[i]]=true
		magicSet[downLsit[i]]=true
	}

	// fmt.Println(upList)
	// fmt.Println(upValNums)
	// fmt.Println(downLsit)
	// fmt.Println(downValNums)
	// fmt.Println(same)
	// fmt.Println(magicSet)
	// fmt.Println(mid)
	flag:=false
	res:=math.MaxInt
	//遍历魔法集合
	for magic, _ := range magicSet {
		if upValNums[magic]+same[magic]>=mid{
			flag=true
			res=0
		}
		if upValNums[magic]+downValNums[magic]+same[magic]>=mid{ //可以组成
			flag=true
			temp:=downValNums[magic] 
			res=min(res,temp)
		}
	}
	if flag==false{
		fmt.Println(-1)
	}else{
		fmt.Println(res)
	}
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
func help(a int)int  {
	if a%2==0{
		return a/2
	}else{
		return a/2+1
	}
}