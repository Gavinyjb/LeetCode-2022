package Problem_0050_PowXN

func myPow(x float64, n int) float64 {
	isNeg := false //正负判断
	if n < 0 {
		isNeg = true //为负值
		n = -n
	} else if n == 0 {
		return 1.0
	}
	i := 2
	res := x
	myMap := make(map[int]float64)
	myMap[1] = x
	for i <= n {
		res = res * res
		myMap[i] = res
		i = i * 2
	}
	i = i / 2
	for k := i / 2; k > 0; k = k / 2 {
		if i == n {
			break
		}
		if k+i <= n {
			res = res * myMap[k]
			i = i + k
		}
	}
	if !isNeg {
		return res
	} else {
		return 1 / res
	}
}
