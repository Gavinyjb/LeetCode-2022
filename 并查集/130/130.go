package leetcode

type UnionFind struct {
	parents []int
	rank    []int
	count   int
}

func (uf *UnionFind) Init(n int) {
	uf.parents = make([]int, n)
	uf.rank = make([]int, n)
	uf.count = 0
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
}
func (uf *UnionFind) Find(q int) int {
	root := q
	for uf.parents[root] != root {
		root = uf.parents[root]
	}

	//路径压缩
	for uf.parents[q] != q {
		temp := uf.parents[q]
		uf.parents[q] = root
		q = temp
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
		uf.parents[proot] = qroot
	} else {
		uf.parents[qroot] = proot
		if uf.rank[qroot] == uf.rank[proot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}
func (uf *UnionFind) GetCount() int {
	return uf.count
}
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board[0]), len(board)
	uf := new(UnionFind)
	uf.Init(n*m + 1) // 特意多一个特殊点用来标记

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i == 0 || i == n-1 || j == 0 || j == m-1) && board[i][j] == 'O' { //棋盘边缘上的 'O' 点
				uf.Union(i*m+j, n*m)
			} else if board[i][j] == 'O' { //棋盘非边缘上的内部的 'O' 点
				if board[i-1][j] == 'O' {
					uf.Union(i*m+j, (i-1)*m+j)
				}
				if board[i+1][j] == 'O' {
					uf.Union(i*m+j, (i+1)*m+j)
				}
				if board[i][j-1] == 'O' {
					uf.Union(i*m+j, i*m+j-1)
				}
				if board[i][j+1] == 'O' {
					uf.Union(i*m+j, i*m+j+1)
				}

			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if uf.Find(i*m+j) != uf.Find(n*m) {
				board[i][j] = 'X'
			}
		}
	}
}
