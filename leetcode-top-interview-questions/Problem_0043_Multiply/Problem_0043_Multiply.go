package Problem_0043_Multiply

func multiply(nums1 string, nums2 string) string {
	if nums1 == "0" || nums2 == "0" {
		return "0"
	}
	ans := "0"
	_, n := len(nums1), len(nums2)
	jinwei := 0 //进位
	for j := n - 1; j >= 0; j-- {

		b := int(nums2[j] - '0')
		one := "0"
		//计算有b个nums1
		for k := 0; k < b; k++ {
			one = plus(one, nums1)
		}
		//计算后面需要加几个“0” 因为进位
		for k := 0; k < jinwei; k++ {
			one = one + "0"
		}
		jinwei++ //每次循环＋1
		ans = plus(ans, one)
	}
	return ans
}

func plus(nums1 string, nums2 string) string {
	num1 := []byte(nums1)
	num2 := []byte(nums2)
	ans := make([]byte, 0)
	i, j := len(num1)-1, len(num2)-1
	flag := byte(0)
	k := 0 //ans 下标
	for ; i >= 0 && j >= 0; i, j, k = i-1, j-1, k+1 {
		one := num1[i] - '0'
		two := num2[j] - '0'
		sum := one + two + flag
		if sum > 9 {
			sum = sum - 10
			flag = byte(1)
		} else {
			flag = byte(0)
		}
		ans = append(ans, sum+'0')
	}
	for ; i >= 0; i-- {
		one := num1[i] - '0'
		sum := one + flag
		if sum > 9 {
			sum = sum - 10
			flag = byte(1)
		} else {
			flag = byte(0)
		}
		ans = append(ans, sum+'0')
	}
	for ; j >= 0; j-- {
		two := num2[j] - '0'
		sum := two + flag
		if sum > 9 {
			sum = sum - 10
			flag = byte(1)
		} else {
			flag = byte(0)
		}
		ans = append(ans, sum+'0')
	}
	if flag != byte(0) {
		ans = append(ans, flag+'0')
	}
	return reverse(ans)
}
func reverse(in []byte) string {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return string(in)
}
