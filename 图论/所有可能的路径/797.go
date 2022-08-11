package leetcode

/*
graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边)

输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/all-paths-from-source-to-target
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


func allPathsSourceTarget(graph [][]int) [][]int {
	res:=make([][]int,0)// 0节点到终点的路径
	path:=make([]int,0)// 0节点到终点的路径
	var dfs func([][]int,int)
	path = append(path, 0)
	dfs=func(graph [][]int, x int)  {
		// 要求从节点 0 到节点 n-1 的路径并输出，所以是 len(graph)-1
		if x==len(graph)-1{  // 找到符合条件的一条路径
			temp:=make([]int,len(path))
			copy(temp,path)
			res = append(res, temp)
		}
		for i := 0; i < len(graph[x]); i++ { // 遍历节点n链接的所有节点
			path = append(path, graph[x][i])// 遍历到的节点加入到路径中来
			dfs(graph,graph[x][i]) // 进入下一层递归
			path=path[:len(path)-1]   // 回溯，撤销本节点
		}
	}
	dfs(graph,0)
	return res
}


