package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func max(a,b int)int  {
	if a<b{
		return b
	}
	return a
}
func dp(w,v []int,target int)int{
	if w==nil||v==nil||len(w)!=len(v)||len(w)==0{
		return 0
	}
	N:=len(w)
	//index 0-N
	//rest 0~target
	dp:=make([][]int,N+1)
	//初始化dp
	for i := 0; i < N+1; i++ {
		dp[i]=make([]int, target+1)
	}
	for i := 0; i < N+1; i++ {
		dp[N][i]=0    //dp[N][...]=0
	}

	for index := N-1; index >=0; index-- {
		for rest := 0; rest <=target; rest++ { //每行只会依赖下一行
			p1:=dp[index+1][rest]
			p2:=0
			next:=-1
			if rest-w[index]<0{
				next=-1
			}else {
				next=dp[index+1][rest-w[index]]
			}
			if next!=-1{
				p2=v[index]+next
			}
			dp[index][rest]=max(p1,p2)
		}
	}
	return dp[0][target]
	
}
func main() {
	v := make([]int, 0) //工作量
	w := make([]int, 0) //重要性
	var input string
	
	inputreader:=bufio.NewReader(os.Stdin)
	input,_=inputreader.ReadString('\n')
	fmt.Println(input)
	vws:=strings.Split(input," ")
	for i := 0; i < len(vws); i++ {
		vw:=strings.Split(vws[i],",") //"v,w"
		v1,_:=strconv.Atoi(vw[0])
		w1,_:=strconv.Atoi(vw[1])
		v = append(v, v1)
		w = append(w, w1)
	}

	fmt.Println(v)
	fmt.Println(w)
	target:=20
	res:=dp(v,w,target)
	fmt.Println(res)
}