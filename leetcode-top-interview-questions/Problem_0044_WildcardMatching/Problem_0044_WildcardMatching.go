package Problem_0044_WildcardMatching

func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}
	s2Byte := []byte(s)
	p2Byte := []byte(p)
	return dp(s2Byte, p2Byte)
}
func dp(s, p []byte) bool {
	//x:->i+1 y:->j+1
	dpMap := make([][]bool, len(s)+1)
	for i, _ := range dpMap {
		dpMap[i] = make([]bool, len(p)+1)
	}
	dpMap[0][0] = true
	//if j==-1&&i==-1{
	//	return true
	//}
	for y := 1; y < len(p)+1; y++ {
		if p[y-1] == '*' {
			dpMap[0][y] = dpMap[0][y-1]
		} else {
			dpMap[0][y] = false
		}
	}
	//if i==-1{ //s 已经为空
	//	if p[j]=='*'{
	//		return try(s,p,i,j-1)
	//	}else {
	//		return false
	//	}
	//}
	for x := 1; x < len(s)+1; x++ {
		dpMap[x][0] = false
	}
	//if j==-1&&i>-1{  //p已经匹配 s还有
	//	return false
	//}
	for x := 1; x < len(s)+1; x++ {
		for y := 1; y < len(p)+1; y++ {
			if s[x-1] == p[y-1] || p[y-1] == '?' { // i j 匹配
				dpMap[x][y] = dpMap[x-1][y-1]
			} else {
				if p[y-1] == '*' {
					dpMap[x][y] = dpMap[x-1][y] || dpMap[x][y-1] //*匹配一位或者0位
				} else {
					dpMap[x][y] = false
				}
			}
		}
	}
	return dpMap[len(s)][len(p)]
}

//回溯法
func isMatch2(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}
	s2Byte := []byte(s)
	p2Byte := []byte(p)
	return try(s2Byte, p2Byte, len(s)-1, len(p)-1)
}
func try(s, p []byte, i, j int) bool {
	if j == -1 && i == -1 {
		return true
	}
	if i == -1 { //s 已经为空
		if p[j] == '*' {
			return try(s, p, i, j-1)
		} else {
			return false
		}
	}
	if j == -1 && i > -1 { //p已经匹配 s还有
		return false
	}
	if s[i] == p[j] || p[j] == '?' { // i j 匹配
		return try(s, p, i-1, j-1)
	} else {
		if p[j] == '*' {
			return try(s, p, i-1, j) || try(s, p, i, j-1) //*匹配一位或者0位
		} else {
			return false
		}
	}
}
