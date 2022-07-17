package leetcdoe

import (
	"fmt"
)

type UnionFind struct {
	parents []int
	rank    []int
	count   int
}

func (uf *UnionFind) Init(n int) {
	uf.parents = make([]int, n)
	uf.rank = make([]int, n)
	uf.count = n
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
}

func (uf *UnionFind) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parents[proot] = qroot
	} else {
		uf.parents[qroot] = proot
		if uf.rank[qroot] == uf.rank[proot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}

//func (uf *UnionFind) Find(p int) int {
//	for uf.parents[p] != p {
//		uf.parents[p] = uf.Find(uf.parents[p])
//	}
//	return uf.parents[p]
//}
func (uf *UnionFind) Find(x int) int {
	root := x
	for uf.parents[root] != root {
		root = uf.parents[root]
	}
	//路径压缩
	for uf.parents[x] != x {
		temp := uf.parents[x]
		uf.parents[x] = root
		x = temp
	}
	return root
}
func (uf *UnionFind) Connect(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}
func (uf *UnionFind) GetCount() int {
	return uf.count
}

func equationsPossible(equations []string) bool {
	if len(equations) == 0 {
		return false
	}
	uf := new(UnionFind)
	uf.Init(26)
	for _, equation := range equations {
		if equation[1] == '=' {
			uf.Union(int(equation[0]-'a'), int(equation[3]-'a'))
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			if uf.Connect(int(equation[0]-'a'), int(equation[3]-'a')) {
				return false
			}
		}
	}
	return true
}
func main() {
	equations := []string{"a==b", "c==d", "a==c", "a!=d"}
	ret := equationsPossible(equations)
	fmt.Println(ret)
}
