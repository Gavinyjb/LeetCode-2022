package main

import "fmt"

type UnionFindSet struct {
	parent map[int]int
	rank   map[int]int
	count  int
}

func (uf *UnionFindSet) Init(n int) {
	uf.parent = make(map[int]int)
	uf.rank = make(map[int]int)
	uf.count = n
	for i := 1; i <= n; i++ {
		one := i
		uf.parent[one] = one
		uf.rank[one] = 0
	}
}

func (uf *UnionFindSet) Find(q int) int {
	root := q
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	//路径压缩
	for uf.parent[q] != root {
		temp := uf.parent[q]
		uf.parent[q] = root
		q = temp
	}
	return root
}

func (uf *UnionFindSet) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.parent[pRoot] = qRoot
	} else {
		uf.parent[qRoot] = pRoot
		if uf.rank[pRoot] == uf.rank[qRoot] {
			uf.rank[pRoot]++
		}
	}
	uf.count--
}
func (uf *UnionFindSet) IsSameSet(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
	// return uf.parent[p]==uf.parent[q]
}
func main() {
	n, m := 0, 0
	fmt.Scanf("%v %v\n",&n,&m)
	var myUF UnionFindSet
	myUF.Init(n)
	for i := 0; i < m; i++ {
		opt,p,q:=0,0,0
		fmt.Scanf("%v %v %v\n",&opt,&p,&q)
		if opt==1 {
			if myUF.IsSameSet(p,q){
				fmt.Println("Yes")
			}else{
				fmt.Println("No")
			}
		}else if opt==2{
			myUF.Union(p,q)
		}
	}
}
