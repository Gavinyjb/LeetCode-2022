package leetcode
type Node struct{
	num int
	Nexts []*Node
}
type Graph struct{
	Nodes map[int]*Node
}
func creatGraph(numCourses int, prerequisites [][]int)*Graph{
	graph:=new(Graph)
	graph.Nodes=make(map[int]*Node)
	for i := 0; i < numCourses; i++ {
		graph.Nodes[i]=&Node{
			num: i,
		}
	}
	for _, v := range prerequisites {
		fromNum:=v[0]
		toNum:=v[1]
		from:=graph.Nodes[fromNum]
		to:=graph.Nodes[toNum]
		from.Nexts = append(from.Nexts, to)
	}
	return graph
}


func canFinish(numCourses int, prerequisites [][]int) bool {
	graph:=creatGraph(numCourses,prerequisites)
	// 记录一次递归堆栈中的节点
	onPath :=make([]bool, numCourses)
	// 记录遍历过的节点，防止走回头路
	visited :=make([]bool, numCourses)
	// 记录图中是否有环
	hasCycle:= false

	var traverse func (*Graph , int)
	traverse=func (graph *Graph ,s int)  {
		if onPath[s]{
			hasCycle=true //出现环路
		}
		if visited[s]||hasCycle{
			// 如果已经找到了环，也不用再遍历了
			return
		}
	
		visited[s]=true
		onPath[s]=true
		for _, t := range graph.Nodes[s].Nexts {
			traverse(graph,t.num)
		}
		onPath[s]=false
	}
	for i := 0; i < numCourses; i++ {
		 // 遍历图中的所有节点
		 traverse(graph, i)
	}
	// 只要没有循环依赖可以完成所有课程
	return !hasCycle
}

// func traverse(graph *Graph ,s int)  {
// 	if onPath[s]{
// 		hasCycle=true //出现环路
// 	}
// 	if visited[s]||hasCycle{
// 		// 如果已经找到了环，也不用再遍历了
// 		return
// 	}

// 	visited[s]=true
// 	onPath[s]=true
// 	for _, t := range graph.Nodes[s].Nexts {
// 		traverse(graph,t.num)
// 	}
// 	onPath[s]=false

// }