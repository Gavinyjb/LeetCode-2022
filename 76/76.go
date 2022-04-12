package main

import "fmt"

func minWindow(s string, t string) string {
	if len(t)>len(s){
		return ""
	}
	if s==t{
		return s
	}
	left,right,res:=0,0,s+" "
	//建立一个哈希表，记录已经已经匹配的t字符串
	hashMapT:=make(map[byte]int,0)
	hashMapW:=make(map[byte]int,0)
	for i := 0; i < len(t); i++ {
		hashMapT[t[i]]++
	}
	for ;right<len(s);right++{
		if hashMapT[s[right]]>0{
			hashMapW[s[right]]++
		}
		for left<=right&&isContain(hashMapW,hashMapT){
			if len(res)>right-left+1{
				res=s[left:right+1]
			}
			if hashMapT[s[left]]==0{
				left++
			}else {
				hashMapW[s[left]]--
				left++
			}
		}
	}
	if res==s+" "{
		return ""
	}
	return res
}
func isContain(w,t map[byte]int)(is bool){
	for k, v := range t {
		if w[k]<v{
			return false
		}
	}
	return true
}
func main() {
	s := "a"
	t := "b"
	ret:=minWindow(s,t)
	fmt.Println(ret)

}
