//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
////type Bridges struct {
////	//任务列表
////	taskNumber []int  //eg：[1 2 3 4 5]
////
////	//场桥列表
////	fieldBridge []int  //eg：[1 2 2 1 1]
////}
//////随机交换
////func (demo *Bridges) randomExchange()  {
////
////}
//////随机交换反序
////func (demo *Bridges) reRandomExchange()  {
////
////}
////
//////逆序
////func (demo *Bridges) reverseOrder()  {
////
////}
////
////func (demo *Bridges) judge() bool {
////	return false
////}
//var(
//	i int //初始容量
//)
////初始化种群
//func initial (nums int)[]int{
//	individual:=make( []int,nums)
//	for i := 0; i < nums; i++ {
//		individual[i]=i+1
//	}
//	fmt.Println(individual)
//	return individual
//}
////交换
//func swap(individual []int) [][]int {
//	res:=make([][]int,0)
//	for i, _ := range individual {
//		for j, _ := range individual {
//			temp:=make([]int,len(individual))
//			copy(temp,individual)
//			if i!=j{
//				temp[i],temp[j]=temp[j],temp[i]
//				//fmt.Printf("交换%v和%v:--序列为：%v\n",i+1,j+1,temp)
//				res = append(res, temp)
//			}
//		}
//	}
//	return res
//}
//
////permute 全排列
//func permute1(nums []int) [][]int {
//	if len(nums) == 1 {
//		return [][]int{nums}
//	}
//	var ret [][]int
//	for i := 0; i < len(nums); i++ {
//		buf := make([]int, len(nums)-1)
//		copy(buf, nums[0:i])
//		copy(buf[i:], nums[i+1:])
//		r := permute2(buf)
//		for _, v := range r {
//			ret = append(ret, append(v, nums[i]))
//		}
//	}
//	return ret
//}
//func permute2(nums []int) [][]int {
//	res:=make([][]int,0)
//	track:=make([]int,0)
//	used:=make([]bool,len(nums))
//	var backtrack func(nums []int,track []int,used []bool)
//	backtrack=func (nums []int,track []int,used []bool){
//		// 触发结束条件
//		if len(track)==len(nums){
//			res=append(res,track)
//			return
//		}
//		for i:=0;i<len(nums);i++{
//			// 排除不合法的选择
//			if used[i]==true{ // nums[i] 已经在 track 中，跳过
//				continue
//			}else{
//				// 做选择
//				track=append(track,nums[i])
//				used[i]=true
//				// 进入下一层决策树
//				backtrack(nums,track,used)
//				// 取消选择
//				track=track[:len(track)-1]
//				used[i]=false
//			}
//		}
//	}
//	backtrack(nums,track,used)
//	return res
//}
////随机抽i个个体
//func lot(population [][]int,i int ) (simple [][]int ) {
//	rand.Seed(time.Now().UnixNano())
//	set:=make(map[int]bool,0)
//	fmt.Println(set)
//	for ;len(set)<i; {
//		j:=rand.Intn(len(population))
//		if set[j]!=true{
//			simple = append(simple, population[j])
//		}
//		set[j]=true
//		fmt.Println(set)
//		fmt.Println(simple)
//	}
//	return simple
//}
////随机插入
//func randInsert(individual []int,i int,j int)[]int  {
//	if i>j{
//		i,j=j,i
//	}
//	buf:=make([]int,len(individual))
//	copy(buf,individual[0:i])
//	copy(buf[i:j],individual[i+1:j+1])
//	buf[j]=individual[i]
//	copy(buf[j+1:],individual[j+1:])
//	return buf
//}
//
//func reverse(individual []int)[]int  {
//	for i,j:=0,len(individual)-1; i <j ;  {
//		individual[i],individual[j]=individual[j],individual[i]
//		i++
//		j--
//	}
//	return individual
//}
//
//func reversePart(individual []int ,index1 int,index2 int)[]int{
//	//buf:=make([]int,len(individual))
//	//copy(buf,individual[:index1])
//	reverse(individual[index1:index2+1])
//
//	return individual
//}
//
//
//
//func main() {
//	individual:=[]int{1,2,3,4}
//	//res:=swap(individual)
//	//for i2, re := range res {
//	//	fmt.Printf("%v:--%v\n",i2,re)
//	//}
//	res2:=permute2(individual)
//	for i2, ints := range res2 {
//		fmt.Printf("%v:--%v\n",i2,ints)
//	}
//	//res3:=lot(res2,5)
//	//for i2, ints := range res3 {
//	//	fmt.Printf("%v:--%v\n",i2,ints)
//	//}
//	//rand.Seed(time.Now().UnixNano())
//	//index1 :=rand.Intn(len(individual))
//	//index2:=rand.Intn(len(individual))
//	//if index1 !=index2{
//	//	res4:=randInsert(individual, index1,index2)
//	//	fmt.Printf("%v插入%v:---%v\n", index1,index2,res4)
//	//}else{
//	//	for{
//	//		index2=rand.Intn(len(individual))
//	//		if index1 !=index2{
//	//			break
//	//		}
//	//	}
//	//	res4:=randInsert(individual, index1,index2)
//	//	fmt.Printf("%v插入%v:---%v\n", index1,index2,res4)
//	//}
//
//	//res5:=reverse(individual)
//	//fmt.Println(res5)
//
//	//res6:=reversePart(individual,2,4)
//	//fmt.Println(res6)
//	//list:=[]int{1,2,3,4,5}
//	//for index1 := 0; index1 < len(list); index1++ {
//	//	for j:=0;j<len(list);j++ {
//	//		temp:=make([]int,5)
//	//		copy(temp,list)
//	//		if index1!=j{
//	//			temp[index1],temp[j]=temp[j],temp[index1]
//	//			fmt.Printf("交换%v和%v:--序列为：%v\n",index1+1,j+1,temp)
//	//		}
//	//	}
//	//}
//}
