package problem0020validparentheses


func isValid(s string) bool {
	stack:=make([]byte,0)

	for i := 0; i < len(s); i++ {
		if s[i]=='('||s[i]=='{'||s[i]=='['{
			stack = append(stack, s[i])
		}
		if s[i]==')'{
			if len(stack)>0&&stack[len(stack)-1]=='('{
				stack=stack[:len(stack)-1]
			}else {
				return false
			}
		}
		if s[i]=='}'{
			if len(stack)>0&&stack[len(stack)-1]=='{'{
				stack=stack[:len(stack)-1]
			}else {
				return false
			}
		}
		if s[i]==']'{
			if len(stack)>0&&stack[len(stack)-1]=='['{
				stack=stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	
	return len(stack)==0
}
