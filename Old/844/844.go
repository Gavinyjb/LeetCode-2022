package main

/*
示例 1：

输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。
示例 2：

输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 ""。
示例 3：

输入：s = "a#c", t = "b"
输出：false
解释：s 会变成 "c"，但 t 仍然是 "b"。

 */
func backspaceCompare(s string, t string) bool {
	if s==t||(s==""&&t==""){
		return true
	}
	sSlice:=[]byte(s)
	tSlice:=[]byte(t)
	return remove(sSlice)==remove((tSlice))
}
func remove(str []byte) string {
	slow,fast:=0,0
	for ;fast<len(str);fast++{
		if str[fast]=='#'{
			slow=max(0,slow-1)
		}else {
			str[slow]=str[fast]
			slow++
		}
	}
	//fmt.Println(string(str[:slow]))
	return string(str[:slow])
}
func max(a,b int)int  {
	if a>b{
		return a
	}
	return b
}
func main() {
	remove([]byte("ab#c"))
}
