package main

import (
	"fmt"
	"math/rand"
)

/*
大根堆
*/
//下标start开始，到end，有count个数字是解锁的
// type Elem struct{
// 	start int
// 	end   int
// 	count int
// 	index int //堆中的索引
// }
// func GetElem(idx ,k int,counts []int)*Elem{
// 	return &Elem{
// 		start: idx,
// 		end: idx+k-1,
// 		count: counts[idx],
// 		index: idx,
// 	}
// }

// type MyHeap []*Elem
// func InitMyHeap(counts []int,k int)[]*Elem  {
// 	ret:=make([]*Elem,0)
// 	for i := 0; i < len(counts); i++ {
// 		elem:=&Elem{
// 			start: i,
// 			end: i+k-1,
// 			count: counts[i],
// 			index: i,
// 		}
// 		ret = append(ret, elem)
// 	}
// 	return ret
// }
// func (h MyHeap)Less(i,j int)bool  {
// 	return h[i].count>h[j].count
// }
// func (h MyHeap)Len()int  {
// 	return len(h)
// }
// func (h MyHeap)Swap(i ,j int)  {
// 	h[i],h[j]=h[j],h[i]
// 	h[i].index=i
// 	h[j].index=j
// }
// func (h *MyHeap)Push(x interface{}){
// 	n:=len(*h)
// 	elem:=x.(*Elem)
// 	elem.index=n
// 	*h=append(*h, elem)
// }
// func (h *MyHeap)Pop()interface{}{
// 	old:=*h
// 	n:=len(old)
// 	elem:=old[n-1]
// 	old[n-1]=nil
// 	elem.index=-1
// 	*h=old[:n-1]
// 	return elem
// }
// func (h *MyHeap)Update(elem *Elem)  {
// 	elem.count++
// 	heap.Fix(h,elem.index)
// }
//获取一个数组，该数组表述自己以及接下来k-1位上解锁数字的个数
func GetCount(str []byte,k int,nums1ToIdxs map[int]map[int]bool)[]int{
	if len(str)<k{
		return nil
	}
	ret:=make([]int,len(str)-k+1)
	for i := 0; i < len(str)-k+1; i++ { //5-4+1
		for j := 0; j < k; j++ {
			ret[i]+=help(str[j+i])
		}
		if _,ok:=nums1ToIdxs[ret[i]];!ok{ //key=ret[i] 不存在
			nums1ToIdxs[ret[i]]=make(map[int]bool)
			nums1ToIdxs[ret[i]][i]=true
		}else {
			nums1ToIdxs[ret[i]][i]=true
		}
	}
	return ret
}
func help(q byte)int{
	if q=='1'{
		return 1
	}
	return 0
}
func FindZero(start ,k int,str []byte)int  {
	for i := 0; i < k; i++ {
		if str[i+start]=='0'{
			str[i+start]='1'
			return i+start
		}
	}
	return -1
}
func main()  {
	n,k:=0,0
	fmt.Scanf("%v %v\n",&n,&k)
	strSrc:=""
	fmt.Scan(&strSrc)
	str:=[]byte(strSrc)
	fmt.Println(string(str))

	// n,k:=100000,3
	// str:=make([]byte,n)
	// str[99998]='1'
	// str[99999]='1'
	// for i := 0; i < 99997; i++ {
	// 	str[i]='0'
	// }

	result:=make([][]int,0)
	nums1ToIdxs:=make(map[int]map[int]bool) //key:=nums1 val:=counts[]切片的下标集合

	counts:=GetCount(str,k,nums1ToIdxs)
	// fmt.Println(counts)/////////////////////////////////

	for len(nums1ToIdxs[k-1])>0{
		l:=len(nums1ToIdxs[k-1])

		keys:=make([]int,0)
		for k := range nums1ToIdxs[k-1] {
			keys = append(keys, k)
		}

		for i := 0; i < l; i++ {
			startToEnd:=make([]int,2)
			startToEnd[0]=keys[0]
			delete(nums1ToIdxs[k-1],startToEnd[0])//弹出
			keys=keys[1:] //弹出
			startToEnd[1]=startToEnd[0]+k-1
			idx:=FindZero(startToEnd[0],k,str)
			
			for j := 0; j < k; j++ {
				if idx-j<0||idx-j>=len(str)-k+1{
					continue
				}
				// fmt.Print("idx-j:=")
				// fmt.Println(idx-j)
				delete(nums1ToIdxs[counts[idx-j]],idx-j)
				counts[idx-j]++
				if _,ok:=nums1ToIdxs[counts[idx-j]];!ok{
					nums1ToIdxs[counts[idx-j]]=make(map[int]bool)
					nums1ToIdxs[counts[idx-j]][idx-j]=true
				}else {
					nums1ToIdxs[counts[idx-j]][idx-j]=true
				}
			}
			result = append(result, startToEnd)
		}
	}
	for _, v := range str {
		if v=='0'{
			fmt.Println(-1)
		}
	}
	// fmt.Println(str[99997])
	// fmt.Println(str[99998])
	// fmt.Println(str[99999])
	fmt.Println(result)
	fmt.Println(len(result))
}
