package leetcodetopinterviewquestions

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	sToBytes := []byte(s)
	m := make(map[byte]int) //存储byte出现的上一个索引坐标

	length := 1
	l := 0
	for i := 0; i < len(sToBytes); i++ {
		one := sToBytes[i]
		if _, ok := m[one]; !ok { //之前并不存在与map
			m[one] = i
			length = max(length, i-l+1)
			// showmap(m)
		} else {
			// 说明之前已经出现过
			l = max(m[one] + 1,l)
			m[one] = i
			length = max(length, i-l+1)
			// showmap(m)
		}
		// fmt.Printf("l:%v,i:=%v,legth:=%v\n",l,i,length)
	}
	
	return length
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func showmap(m map[byte]int)  {
	for k, v := range m {
		fmt.Printf("%c :%v\n",k,v)
	}
}
