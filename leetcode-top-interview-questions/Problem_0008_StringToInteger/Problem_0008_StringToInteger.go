package leetcodetopinterviewquestions

import (
	"bytes"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s=strings.Trim(s," ")  //去除空格
	
	if len(s)==0{
		return 0
	}
	byteS := []byte(s)
	var neg bool  //正负判断
	if byteS[0]=='-'{
		neg=false
		byteS=byteS[1:]
	}else if byteS[0]=='+'{
		neg=true
		byteS=byteS[1:]
	}else if byteS[0]<='9'&&byteS[0]>='0'{
		neg=true
	}else { //前缀不合法
		return 0
	}
	res:=help(byteS,neg)
	return res
}

func help(byteS []byte,neg bool)int  {
	//按照负数来做
	res:=0
	minq:=math.MinInt32/10
	minR:=math.MinInt32%10
	for i := 0; i < len(byteS); i++ {
		if byteS[i]<'0'||byteS[i]>'9'{
			break
		}else {
			cur:=int('0')-int(byteS[i])
			if res<minq||(res==minq&&cur<minR){ //必然越界
				if neg{
					return math.MaxInt32
				}else {
					return math.MinInt32
				}
			}
			res=res*10+cur
		}
	}
	if neg{
		return -res
	}else {
		return res
	}
}