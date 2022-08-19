package roblem0028implementstrstr

func strStr(haystack string, needle string) int {

}

func getNextArray(ms []byte )[]int  {
	if len(ms)==1{
		return []int{-1}
	}
	next:=make([]int,len(ms))
	next[0]=-1
	next[1]=0
	i:=2
	
}

// func strStr(haystack, needle string) int {
//     n, m := len(haystack), len(needle)
//     if m == 0 {
//         return 0
//     }
//     pi := make([]int, m)
//     for i, j := 1, 0; i < m; i++ {
//         for j > 0 && needle[i] != needle[j] {
//             j = pi[j-1]
//         }
//         if needle[i] == needle[j] {
//             j++
//         }
//         pi[i] = j
//     }
//     for i, j := 0, 0; i < n; i++ {
//         for j > 0 && haystack[i] != needle[j] {
//             j = pi[j-1]
//         }
//         if haystack[i] == needle[j] {
//             j++
//         }
//         if j == m {
//             return i - m + 1
//         }
//     }
//     return -1
// }