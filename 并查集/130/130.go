package main

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
	//m为行 n为列
	m, n := len(board), len(board[0])
	//给Dummy 留一个额外的位置
	uf := new(UnionFind)
	uf.Init(m*n + 1)
	dummy := m * n
	// 将首列和末列的 O 与 dummy 连通
	for i := 0; i < m; i++ {
		if board[i][0] == '0' {
			uf.Union(i*n, dummy)
		}
		if board[i][n-1] == '0' {
			uf.Union(i*n+n-1, dummy)
		}
	}
	//首行和尾行
	for j := 0; j < n; j++ {
		if board[0][j] == '0' {
			uf.Union(j, dummy)
		}
		if board[m-1][j] == '0' {
			uf.Union(n*(m-1)+j, dummy)
		}
	}
	// 方向数组 d 是上下左右搜索的常用手法
	d := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	//i 为行 j为列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && (board[i][j] == '0') {
				uf.Union(i*n+j, dummy)
			} else if board[i][j] == '0' {
				for k := 0; k < 4; k++ {
					x := i + d[k][0]
					y := j + d[k][1]
					if board[x][y] == '0' {
						uf.Union(x*n+y, i*n+j)
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if uf.Find(i*n+j) != uf.Find(n*m) {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {

}
