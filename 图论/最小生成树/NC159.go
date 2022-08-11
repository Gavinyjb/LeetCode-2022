package main

import "container/heap"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 返回最小的花费代价使得这n户人家连接起来
 * @param n int n户人家的村庄
 * @param m int m条路
 * @param cost int二维数组 一维3个参数，表示连接1个村庄到另外1个村庄的花费的代价
 * @return int
 */
type Node struct{
	Val int
	In  int //入度
	Out int //出度
	Nexts []*Node  //邻居节点
	Edge  []*Edge  //所出的边
}
func InitNode(val int)*Node  {
	return &Node{
		Val: val,
		In: 0,
		Out: 0,
		Nexts: make([]*Node,0 ),
		Edge: make([]*Edge, 0),
	}
}
type Edge struct{
	Wight int //边权重
	From *Node 
	To *Node 
}
func InitEdge(weight int,from ,to *Node)*Edge  {
	return &Edge{
		Wight: weight,
		From: from,
		To: to,
	}
}
type Graph struct{
	Nodes map[int]*Node
	edges map[*Edge]bool
}

func createGraph(matrix [][]int)Graph{
	greph:=new(Graph)
	greph.Nodes=make(map[int]*Node)
	greph.edges=make(map[*Edge]bool)

	for i := 0; i < len(matrix); i++ {
		//拿到每一条边
		from:=matrix[i][0]
		to:=matrix[i][1]
		weight:=matrix[i][2]
		if _,ok:=greph.Nodes[from];!ok{
			greph.Nodes[from]=InitNode(from)
		}
		if _,ok:=greph.Nodes[to];!ok{
			greph.Nodes[to]=InitNode(to)
		}
		fromNode:=greph.Nodes[from]
		toNode:=greph.Nodes[to]

		newEdge:=InitEdge(weight,fromNode,toNode)
		fromNode.Nexts = append(fromNode.Nexts, toNode)
		fromNode.Out++
		toNode.In++
		fromNode.Edge = append(fromNode.Edge, newEdge)
		greph.edges[newEdge]=true
	}
	return *greph
}
/*
并查集
*/
type UnionFind struct{
	// key 某一个节点， value key节点往上的节点
	fathermap map[*Node]*Node
	// key 某一个集合的代表节点, value key所在集合的节点个数
	rank map[*Node]int
}
//获取UnionFind
func GetUnionFind()UnionFind{
	return UnionFind{
		fathermap: make(map[*Node]*Node),
		rank: make(map[*Node]int),
	}
}
// 从Node切片初始化UnionFind
func (uf *UnionFind)Init(nodes []*Node)  {
	for _,v := range nodes {
		uf.fathermap[v]=v
		uf.rank[v]=1
	}
}
func (uf *UnionFind)Find(q *Node)*Node  {
	path:=make([]*Node,0)
	for uf.fathermap[q]!=q{
		path = append(path, q)
		q=uf.fathermap[q]
	}
	for len(path)>0{
		node:=path[0]
		path=path[1:]
		uf.fathermap[node]=q
	}
	return q
}
func (uf *UnionFind)Union(p,q *Node)  {
	pFather:=uf.fathermap[p]
	qFather:=uf.fathermap[q]
	if qFather==pFather{
		return 
	}
	if uf.rank[pFather]<=uf.rank[qFather]{
		uf.fathermap[pFather]=qFather
		uf.rank[qFather]=uf.rank[qFather]+uf.rank[pFather]
		delete(uf.rank,pFather)
	}else {
		uf.fathermap[qFather]=pFather
		uf.rank[pFather]=uf.rank[qFather]+uf.rank[pFather]
		delete(uf.rank,qFather)
	}
}
func (uf *UnionFind)IsConnect(p,q *Node)bool  {
	return uf.Find(p)==uf.Find(q)
}
/*
小根堆
*/
type EdgeHeap []*Edge
func (h EdgeHeap)Less(i,j int)bool  {
	return h[i].Wight<h[j].Wight
}
func (h EdgeHeap)Len()int  {
	return len(h)
}
func (h EdgeHeap)Swap(i ,j int)  {
	h[i],h[j]=h[j],h[i]
}
func (h *EdgeHeap)Push(x interface{}){
	*h=append(*h, x.(*Edge))
}
func (h *EdgeHeap)Pop()interface{}{
	old:=*h
	x:=old[len(old)-1]
	*h=old[:len(old)-1]
	return x
}
func kruskalMST(greap Graph)[]*Edge{
	uf:=GetUnionFind()
	nodes:=make([]*Node,0)
	for _, v := range greap.Nodes {
		nodes = append(nodes, v)
	}
	
	uf.Init(nodes)
	//从小到大的边，依次弹出，小根堆！！！
	edges:=make([]*Edge,0)
	for v, _ := range greap.edges {
		edges=append(edges, v)
	}
	edgeHeap:=new(EdgeHeap)
	heap.Init(edgeHeap)
	for i := 0; i < len(edges); i++ { //m条边
		heap.Push(edgeHeap,edges[i])  //O log(m)
	}
	ret:=make([]*Edge,0)
	for edgeHeap.Len()>0{
		edge:=heap.Pop(edgeHeap).(*Edge)
		if !uf.IsConnect(edge.From,edge.To){
			ret = append(ret, edge)
			uf.Union(edge.From,edge.To)
		}
	}
	return ret
}
func miniSpanningTree( n int ,  m int ,  cost [][]int ) int {
    // write code here
	greap:=createGraph(cost)
	edges:=kruskalMST(greap)
	res:=0
	for _, v := range edges {
		res+=v.Wight
	}
	return res
}