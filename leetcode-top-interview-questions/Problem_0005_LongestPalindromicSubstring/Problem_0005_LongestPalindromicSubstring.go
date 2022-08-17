package leetcodetopinterviewquestions

import (
	"fmt"
	"math"
)
//manacher算法
func longestPalindrome(s string) string {
	//1)回文半径数组
	//          # 1 # 2 # 3 # 2 # 1 #...M
	//parr[]int{1,2,1,2,1,6...........}

	byteArr := manacherString(s)
	pArr := make([]int, len(byteArr)) //回文半径数组
	index := -1   //最右的扩张成功的位置对应的中心轴下标
	pR := -1     //代表最右的扩张成功的位置
	max := math.MinInt  //结果的长度
	mid:=-1  //结果的中轴
	for i := 0; i < len(byteArr); i++ {
		if i>pR{   //i在R外
			//暴力扩大 pR变大
			pArr[i]=1
		}else {
			//i对称点的回文半径:pArr[2*index-i]
            pArr[i]=min(pArr[2*index-i],pR-i+1)    
			// if (){    //i对称点的回文区域彻底在L...R内部
			// 	pArr[i]=pArr[2*index-i]
			// }else if (){  //i对称点的回文区域跑到L...R外部
			// 	pArr[i]=pR-i+1
			// }else {   //i的对称点的回文区域左边界和L压线
			// 	//从pR开始往外扩张 R变大
			// }
			// !!!!!
			// 以上这些我都不去考虑，直接扩大，但前两种会直接退出，不会进入循环
		}
		for i+pArr[i]<len(byteArr)&&i-pArr[i]>-1{
			if byteArr[i+pArr[i]]==byteArr[i-pArr[i]]{
				pArr[i]++
			}else {
				break
			}
		}
		if i+pArr[i]-1>pR{
			pR=i+pArr[i]-1
			index=i
		}
		if max<pArr[i]{
			max=pArr[i]
			mid=i
		}
	}
	// byteArr:#b#a#b#a#d#
	// mid:3 max:4
	res:=byteArr[mid-max+1:mid+max]  // #b#a#b#
	result:=make([]byte,0)
	for i := 0; i < len(res); i++ {
		if res[i]!='#'{
			result = append(result, res[i])
		}
	}
	return string(result)

}
func manacherString(str string) []byte {
	//str :         abc123321
	//newStrToByte: #a#b#c#1#2#3#3#2#1#
	strToByte := []byte(str)
	newStrToByte := make([]byte, 0)
	for _, v := range strToByte {
		newStrToByte = append(newStrToByte, '#', v)
	}
	newStrToByte = append(newStrToByte, '#')
	return newStrToByte
}
func min(a,b int)int  {
	if a<b{
		return a
	}
	return b
}
//暴力中心扩张 O(n^2)
func longestPalindrome2(str string) string {
	//str : abc123321
	//new : #a#b#c#1#2#3#3#2#1#
	strToByte := []byte(str)
	new := make([]byte, 0)
	for _, v := range strToByte {
		new = append(new, '#', v)
	}
	new = append(new, '#')
	res := 0
	ret := ""
	for i := 0; i < len(new); i++ {
		l, r := i, i
		for l >= 0 && r < len(new) {
			if new[l] == new[r] {
				if r-l+1 > res {
					res = r - l + 1
					ret = string(new[l : r+1])
				}
				l--
				r++
			} else {
				break
			}
		}
	}
	result := make([]byte, 0)
	for i := 0; i < len(ret); i++ {
		if ret[i] != '#' {
			result = append(result, ret[i])
		}
	}
	return string(result)
}
