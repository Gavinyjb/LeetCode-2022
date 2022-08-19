package main

import (
	"fmt"
	"math"
)

// 怪兽还剩hp 滴血 hp初始为100
// 每次的伤害在 0.3 3次砍到 三次没砍到
func process(list [][]int,idx int,rest int,m int, d int,hp int) int {
	if rest==0{  //炮击结束
		if hp<=0{
			return 1
		}else {
			return 0
		}
	}
	// 如果怪兽没有血 || 血量为负
    // 可以直接返回砍死的情况数：myPow(30,rest)
	// if hp<=0{
	// 	return myPow(1000,rest)
	// }
	ways:=0
	//idx 第idx炮击
	ret:=make([]int,0)  //伤害列表
	fmt.Println(ret)
	for i := 0; i < 3; i++ {
		p0:=list[idx][0]
		p1:=list[idx][1]
		p2:=list[idx][2]
		ret=help(p0,p1,p2)
	}
	for i := 0; i < len(ret); i++ {
		heart:=ret[i]*m*d
		ways+=process(list,rest-1,idx+1,m,d,hp-heart)
	}
	return ways
}

func dpFunc(list [][]int,n,m int, d int,hp int) int {
	if n==0{  
		return 0
	}
	 // 行：剩余砍的次数 rest
     // 列：怪兽剩余血量 hp
	 dp:=make([][]int,n+1)
	 for i := 0; i < n+1; i++ {
		temp:=make([]int, 101)
		dp[i]=temp
	 }

	 dp[0][0]=1
	 for times := 1; times <n ; times++ {
		// 填第0列的值 第0列的值：怪兽血量为0，可以直接返回砍死的情况数
		dp[times][0]=myPow(1000,times)
		for hp := 1; hp <=100; hp++ {
			ways:=0
			//idx 第idx炮击
			ret:=make([]int,0)  //伤害列表
			for i := 0; i < 3; i++ {
				p0:=list[times][0]
				p1:=list[times][1]
				p2:=list[times][2]
				ret=help(p0,p1,p2)
			}
			for i := 0; i < len(ret); i++ {
				heart:=ret[i]*m*d
				if hp-heart<=0{
					ways+=dp[times-1][0]
				}else {
					ways+=dp[times-1][hp-heart]
				}
				
			}
			dp[times][hp]=ways
		}
	 }
	 kill:=dp[n][100]
	 fmt.Println(kill)
	 return kill
}
func help(a,b,c int)[]int{
	ret:=make([]int,0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				one:=0
				if i<a{
					one++
				}
				if j<b{
					one++
				}
				if k<c{
					one++
				}
				ret = append(ret, one)
			}
		}
	}
	return ret
}
func main() {
	n, m, d := 0, 0, 0 //n轮炮击  火焰持续m轮  伤害d%
	fmt.Scanf("%d %d %d\n",&n,&m,&d)
	fmt.Println(m,n,d)
	list:=make([][]float64,n)
	for i := 0; i < n; i++ {
		temp:=make([]float64,3)
		fmt.Scanf("%f %f %f\n",&temp[0],&temp[1],&temp[2])
		list[i]=temp
	}
	fmt.Println(list)
	helps:=make([][]int,n)
	for i := 0; i < n; i++ {
        temp:=make([]int,3)
		temp[0]=int(list[i][0])
		temp[1]=int(list[i][1])
		temp[2]=int(list[i][2])
        helps[i]=temp
        
	}
	fmt.Println(helps)
	all:=myPow(1000,n)
	kill:=dpFunc(helps,n,m,d,100)
	result:=float64(kill)/float64(all)
	fmt.Println(result)

}
func myPow(a,b int) int {
	newA:=float64(a)
	newB:=float64(b)
	return int(math.Pow(newA,newB))
}