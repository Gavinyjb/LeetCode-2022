package problem0010regularexpressionmatching



func isMatch2(s string, p string) bool{
	i,j:=len(s)-1,len(p)-1
	s2Bytes:=[]byte(s)
	p2Bytes:=[]byte(p)
	return try(s2Bytes,p2Bytes,i,j)
}

func try(s,p []byte,i,j int)bool  {
	if j==-1&&i>=0{  //p为空串，s不为空串，肯定不匹配。
		return false
	}
	if i==-1&&j==-1{  //s、p都为空串，肯定匹配。
		return true
	}
	if i==-1&&j>0{  //s为空串，但p不为空串，要想匹配，只可能是右端是星号，它干掉一个字符后，把 p 变为空串。
		if p[j]=='*'{
			return try(s,p,i,j-2)
		}else {
			return false
		}
	}
	if s[i]==p[j]||p[j]=='.'{  //如果 匹配
		return try(s,p,i-1,j-1)
	}else {
		if p[j]=='*'{  //p[j]为*
			if p[j-1]==s[i]||p[j-1]=='.'{   //p[j]为*，且p[j-1]与s[i]匹配，消去一个s[i]
				return try(s,p,i-1,j)||try(s,p,i,j-2)        //因为有时候即使匹配也未必一定去匹配  比如：s:"aaa"  p:"ab*a*c*a"
			}else {           //p[j]为*，且p[j-1]与s[i]不匹配,消去  p[j-1] *
				return try(s,p,i,j-2)
			}
		}else{ //不匹配，且最后一位不为*
			return false
		}
	}
}
