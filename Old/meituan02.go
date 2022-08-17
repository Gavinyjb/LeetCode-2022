package main

import "fmt"
//获取加工后的数组
//eg:n=4  原：[0,1,2,3,4]  ---》[2 0 4 1 3]
//eg:n=5  原： [0,1,2,3,4,5]  --》[2 0 4 1 3]
//eg:n=6  原： [0,1,2,3,4,5,6]  ---》[2 5 3 1 4 0]
func help(n int)[]int  {
	inds:=make([]int,n)
	res:=make([]int,0)
	for i := 0; i < n; i++ {
		inds[i]=i
	}
	for len(inds)>2{
		res = append(res, inds[2])
		old0,old1:=inds[0],inds[1]
		inds=inds[3:]
		inds = append(inds, old0,old1)
	}
	res = append(res, inds...)
	return res
}

func main() {
	n := 0
	fmt.Scan(&n)
	nums:=make([]int,n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	helps:=help(n)
	// fmt.Println(nums)
	// fmt.Println(helps)
	result:=make([]int,n)
	//遍历索引数组  i为helps数组的下标，利用它获取nums也就是中数字的值
	// idx存的是该数字原本存在于哪个下标，直接写上去
	for i, idx := range helps {
		result[idx]=nums[i]
	}
	fmt.Println(result)
}