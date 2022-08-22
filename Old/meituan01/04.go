package main

import "fmt"

func main() {
	n, m := 0, 0
	fmt.Scanf("%d %d\n", &n, &m)
	aList := make([]int, n)
	bList := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&aList[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&bList[i])
	}
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, m+1)
	}

}

//func getSame(a, b []int)[]int  {
//	res:=make([]int,0)
//	n:=len(a)
//	m:=len(b)
//	for i ,j:= 0,0; i < n&& j<m;  {
//		if a[i]==b[j]{
//			res = append(res, a[i])
//			i++
//			j++
//		}else {
//			if a[i]<b[j]{
//				i++
//			}else {
//				j++
//			}
//		}
//	}
//	return  res
//}
//func getNotSame(a,same []int) []int {
//	res:=make([]int,0)
//	for i,j:= 0,0; i <len(a)&&j<len(same) ;  {
//		if a[i]==same[j]{
//			i++
//			j++
//		}else {
//			res = append(res, a[i])
//			i++
//		}
//	}
//	return res
//}
//func getSameRes(a,b []int)int  {
//	res:=0
//	n:=len(a)
//
//	for i := 0; i < n; i++ {
//		res+=abs(b[i],a[i])
//	}
//	return res
//}
//func abs(a ,b int)int  {
//	if a<b{
//		return b-a
//	}else {
//		return a-b
//	}
//}
//func main()  {
//	n,m:=0,0
//	fmt.Scanf("%d %d\n",&n,&m)
//	aList:=make([]int,n)
//	bList:=make([]int,n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&aList[i])
//	}
//	for i := 0; i < n; i++ {
//		fmt.Scan(&bList[i])
//	}
//	sort.Ints(aList)
//	sort.Ints(bList)
//
//	res:=0
//	same:=make([]int,0)
//	if n==m{
//		res=getSameRes(aList,bList)
//	}else {
//		if n>m{
//			same=getSame(aList,bList)
//			nRe:=getNotSame(aList,same)
//			mRe:=getNotSame(bList,same)
//			if m>len(same){
//				nRe=nRe[:m-len(same)]
//				res=getSameRes(nRe,mRe)
//			}else if m==len(same){
//				for _, v := range nRe {
//					res+=v
//				}
//			}
//		}else {
//			same=getSame(bList,aList)
//			nRe:=getNotSame(aList,same)
//			mRe:=getNotSame(bList,same)
//			if n>len(same){
//				mRe=mRe[:n-len(same)]
//				res=getSameRes(nRe,mRe)
//			}else if n==len(same){
//				for _, v := range mRe {
//					res+=v
//				}
//			}
//		}
//	}
//	fmt.Println(res)
//}
