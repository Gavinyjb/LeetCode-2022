package leetcodetopinterviewquestions

import (
	"math"
	"strconv"
)

// public static int reverse(int x) {
// 	boolean neg = ((x >>> 31) & 1) == 1;
// 	x = neg ? x : -x;
// 	int m = Integer.MIN_VALUE / 10;
// 	int o = Integer.MIN_VALUE % 10;
// 	int res = 0;
// 	while (x != 0) {
// 		if (res < m || (res == m && x % 10 < o)) {
// 			return 0;
// 		}
// 		res = res * 10 + x % 10;
// 		x /= 10;
// 	}
// 	return neg ? res : Math.abs(res);
// }
func reverse(x int) int {
	//判断正负，负数的范围大于正数
	var neg bool  //true 为正  flase 为负
	neg=x>0
	if neg{
		x=-x
	}
	m:=math.MinInt32/10
	o:=math.MinInt32%10
	res:=0
	for x!=0{  //注意x始终为负数
		if res<m||(res==m&&x%10<o){ //必然越界
			return 0
		}
		res=res*10+x%10
		x=x/10
	}
	if neg{
		return -res
	}else {
		return res
	}
}


//转字符串不可取，注意越界 32位变成0
func reverse2(x int) int {
	if x==0{
		return 0
	}
	xTos:=strconv.Itoa(x)
	var flag bool //true 为正  flase 为负
	if xTos[0]=='-'{
		flag=false
		xTos=xTos[1:]
	}else {
		flag=true
	}
	resBytes:=make([]byte,0)
	for i := len(xTos)-1; i >-1; i-- {
		resBytes = append(resBytes, xTos[i])
		xTos=xTos[:i]
	}
	for {
		if resBytes[0]=='0'{
			resBytes=resBytes[1:]
		}else {
			break
		}
	}
	res,_:=strconv.Atoi(string(resBytes))
	if res>math.MaxInt32||res<math.MinInt32{
		return 0
	}
	if flag{
		return res
	}else {
		return -res
	}
}