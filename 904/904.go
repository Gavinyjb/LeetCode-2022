package main

import "fmt"

/*
你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。

你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：

你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到下一棵树，并继续采摘。
一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。
 */
func totalFruit1(tree []int) int {
	if len(tree) == 0 {
		return 0
	}
	left, right, counter, res, freq := 0, 0, 1, 1, map[int]int{}
	freq[tree[0]]++
	for left < len(tree) {
		if right+1 < len(tree) && ((counter > 0 && tree[right+1] != tree[left]) || (tree[right+1] == tree[left] || freq[tree[right+1]] > 0)) {
			if counter > 0 && tree[right+1] != tree[left] {
				counter--
			}
			right++
			freq[tree[right]]++
		} else {
			if counter == 0 || (counter > 0 && right == len(tree)-1) {
				res = max(res, right-left+1)
			}
			freq[tree[left]]--
			if freq[tree[left]] == 0 {
				counter++
			}
			left++
		}
	}
	return res
}
func max(a,b int)int  {
	if a>b{
		return a
	}
	return b
}

func totalFruit(fruits []int) int {
	left,right,res:=0,0,0
	//建立一个哈希表，记录已经在篮子里的水果个数(区间内元素的种类个数)
	hashMap:=make(map[int]int,0)
	for; right <len(fruits) ;right++ {
		//我们想得到以right为结尾的区间，所以i必须取到
		hashMap[fruits[right]]++
		//当哈希表的长度是3的时候表示区间内元素的种类有3个，不符合题意
		for left<=right &&len(hashMap)>2{
			//left指针向右移动，将left指向的水果数目-1
			hashMap[fruits[left]]--
			if hashMap[fruits[left]]==0{
				//当hashMap[fruits[left]] == 0时表示篮子里没有这种水果了，直接delete
				delete(hashMap,fruits[left])
			}
			left++
		}
		//每次循环之后，我们都得到以right为结尾的符合题意的最长区间
		//通过if选择每次的结果，得到最终的最优解
		res=max(res,right-left+1)
	}
	return res
}
func main() {
	fruits := []int{1,2,1}
	ret:=totalFruit(fruits)
	fmt.Println(ret)
}

