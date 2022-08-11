package leetcdoe

import (
	"fmt"
)
// func help(row,col int)string{
// 	return strconv.Itoa(row)+strconv.Itoa(col)
// }
type state struct{
	x int
	y int
}
func help(x,y int)state{
	return state{x,y}
}

type UnionFind struct {
	parent map[state]state
	rank   map[state]int
	count  int
}

func (uf *UnionFind) Init(grid [][]byte) {
	row:=len(grid)
	col:=len(grid[0])

	uf.parent = make(map[state]state)
	uf.rank = make(map[state]int)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j]=='1'{
				uf.count++
			}
			index:=state{i,j}
			uf.parent[index]=index
			uf.rank[index]=0
		}
	}
}
func (uf *UnionFind) Find(q state)state{
	root:=q
	for uf.parent[root]!=root{
		root=uf.parent[root]
	}
	//路径压缩
	for q!=uf.parent[q]{
		temp:=uf.parent[q]
		uf.parent[q]=root
		q=temp
	}
	return root
}
func (uf *UnionFind) Union(p,q state){
	pRoot:=uf.Find(p)
	qRoot:=uf.Find(q)
	if qRoot==pRoot{
		return
	}
	if uf.rank[pRoot]<uf.rank[qRoot]{
		uf.parent[pRoot]=qRoot
	}else{
		uf.parent[qRoot]=pRoot
		if uf.rank[pRoot]==uf.rank[qRoot]{
			uf.rank[pRoot]++
		}
	}
	uf.count--
}
func (uf *UnionFind)GetCount()int{
	return uf.count
}
func (uf *UnionFind)isSameSet(p,q state)bool{
	return uf.parent[q]==uf.parent[p]
}
func numIslands(grid [][]byte) int {
	if grid==nil{
		return 0
	}
	rows:=len(grid)
	cols:=len(grid[0])
	var unionfindset UnionFind
	unionfindset.Init(grid)

	fmt.Println(unionfindset.count)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j]=='1'{
				if i-1>=0&&grid[i-1][j]=='1'{
					unionfindset.Union(help(i-1,j),help(i,j))
				}
				if i+1<rows&&grid[i+1][j]=='1'{
					unionfindset.Union(help(i+1,j),help(i,j))
				}
				if j-1>=0&&grid[i][j-1]=='1'{
					unionfindset.Union(help(i,j-1),help(i,j))
				}
				if j+1<cols&&grid[i][j+1]=='1'{
					unionfindset.Union(help(i,j+1),help(i,j))
				}
			}
		}
	}
	for k, v := range unionfindset.parent {
		fmt.Printf("k:=%v v:=%v",k,v)
	}
	return unionfindset.count
}
