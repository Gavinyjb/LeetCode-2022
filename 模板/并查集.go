package 模板

type UnionFindInterface interface {
	Init()
	Find()
	Union()
	GetCount()
	IsConnect() bool
}
type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func (uf *UnionFind) Init(n int) {
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	uf.count = n
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
}

func (uf *UnionFind) Find(q int) int {
	root := q
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	for q != uf.parent[q] {
		tmp := uf.parent[q]
		uf.parent[q] = root
		q = tmp
	}
	return root
}
func (uf *UnionFind) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if qroot == proot {
		return
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parent[proot] = qroot
	} else {
		uf.parent[qroot] = proot
		if uf.rank[qroot] == uf.rank[proot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}
func (uf *UnionFind) GetCount() int {
	return uf.count
}
func (uf *UnionFind) IsConnect(p, q int) bool {
	return uf.parent[q] == uf.parent[p]
}
