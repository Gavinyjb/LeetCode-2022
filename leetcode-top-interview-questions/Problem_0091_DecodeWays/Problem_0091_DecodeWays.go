package Problem_0091_DecodeWays

func numDecodings(s string) int {
	//ans := process(s, 0)
	ans := dpFunc(s)
	return ans
}
func dpFunc(s string) int {
	dp := make([]int, len(s)+1)
	n := len(s)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			way := dp[i+1]
			if i+1 == n {
				dp[i] = way
			} else {
				num := int(s[i]-'0')*10 + int(s[i+1]-'0')
				if num <= 26 { //最后两位可以变成一位{
					way += dp[i+2]
					dp[i] = way
				} else {
					dp[i] = way
				}
			}
		}
	}
	return dp[0]
}

//从[0,idx]  有几种答案
func process(s string, idx int) int {
	if idx == len(s) {
		return 1
	}
	//
	if s[idx] == '0' {
		return 0
	}
	way := process(s, idx+1)
	if idx+1 == len(s) {
		return way
	}
	num := int(s[idx]-'0')*10 + int(s[idx+1]-'0')
	if num <= 26 { //最后两位可以变成一位{
		way += process(s, idx+2)
		return way
	} else {
		return way
	}
}
