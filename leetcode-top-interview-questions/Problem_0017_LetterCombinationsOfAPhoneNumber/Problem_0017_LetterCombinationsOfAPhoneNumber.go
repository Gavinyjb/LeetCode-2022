package problem0017lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
    if digits==""{
		return nil
	}
	i2a := make(map[byte]([]byte))
	i2a['2'] = []byte{'a', 'b', 'c'}
	i2a['3'] = []byte{'d', 'e', 'f'}
	i2a['4'] = []byte{'g', 'h', 'i'}
	i2a['5'] = []byte{'j', 'k', 'l'}
	i2a['6'] = []byte{'m', 'n', 'o'}
	i2a['7'] = []byte{'p', 'q', 'r', 's'}
	i2a['8'] = []byte{'t', 'u', 'v'}
	i2a['9'] = []byte{'w', 'x', 'y', 'z'}

	digits2Byte := []byte(digits)
	ans := make([][]byte, 0)
	var process func ( []byte,  int,  []byte, *[][]byte)  
	process = func(digits []byte, i int, path []byte, ans *[][]byte) {
		if i == len(digits) {
			temp := make([]byte, len(path))
			copy(temp, path)
			*ans = append(*ans, temp)
		}else {
			for _, v := range i2a[digits[i]] {
				path [i]=v
				process(digits, i+1, path, ans)
                // path=path[:len(path)-1]  //回撤
			}
		}
	}
	path:=make([]byte,len(digits))
	process(digits2Byte, 0, path,&ans)
	ret:=make([]string,len(ans))
	for i, v := range ans {
		ret[i]=string(v)
	}
	return ret
}
