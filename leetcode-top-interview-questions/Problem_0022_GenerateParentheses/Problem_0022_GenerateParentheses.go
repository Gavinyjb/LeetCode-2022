package problem0022generateparentheses


func generateParenthesis(n int) []string {

	path:=make([]byte,2*n)
	ans:=make([]string,0)

	// 依次在path上填写决定
	// ( ( ) ) ( )....
	// 0 1 2 3 4 5
	// path[0...index-1]决定已经做完了
	// index位置上，( )
	var process func(path []byte,index,leftMinusRight,leftRest int)
	//leftMinusRight 已处理的左括号比右括号多的数量,
	// leftRest 还能填的左括号数量
	process=func(path []byte, index, leftMinusRight, leftRest int) {
		if index==len(path){
			temp:=make([]byte, len(path))
			copy(temp,path)
			ans = append(ans, string(temp))
		}else {
			if leftRest>0{
				path[index]='('
				process(path,index+1,leftMinusRight+1,leftRest-1)
			}
			if leftMinusRight>0{
				path[index]=')'
				process(path,index+1,leftMinusRight-1,leftRest)
			}
		}
	}
	process(path,0,0,n)
	return ans
}

//这个有问题

// func generateParenthesis(n int) []string {
// 	path := make([]byte, 0)
// 	ans := make([]string, 0)
// 	if n==0{
// 		return ans
// 	}
// 	if n==1{
// 		return []string{"()"}
// 	}
// 	var process func([]byte,int,int)
// 	process=func (path []byte,index int,n int){
// 		if index==n{
// 			temp:=make([]byte, len(path))
// 			copy(temp,path)
// 			ans = append(ans, string(temp))
// 		}else {
//              m:=make(map[string]int)
// 			for i := 0; i < 3; i++ {  //path ()
// 				if i==0{       //1) ()加到path 后面 ()()
// 					path = append(path, '(',')')
// 					if _,ok:=m[string(path)];!ok{
// 						m[string(path)]=1
// 						process(path,index+1,n)
// 						path=path[:len(path)-2]
// 					}else {
// 						path=path[:len(path)-2]
// 					}
// 				}else if i==1 {
// 					temp:=[]byte{'(',')'}  //2) ()加到path 前面 ()()
// 					path = append(temp,path... )
// 					if _,ok:=m[string(path)];!ok{
// 						m[string(path)]=1
// 						process(path,index+1,n)
// 						path=path[2:]
// 					}else {
// 						path=path[2:]
// 					}
// 				}else {
// 					temp:=[]byte{'('}    //3) ()加到path 两边 (())
// 					path = append(temp,path... )
// 					path = append(path, ')')
// 					if _,ok:=m[string(path)];!ok{
// 						m[string(path)]=1
// 						process(path,index+1,n)
// 						path=path[1:len(path)-1]
// 					}else {
// 						path=path[1:len(path)-1]
// 					}
// 				}
// 			}
// 		}
		
// 	}
// 	process(path,0,n)
// 	return ans
// }


   



