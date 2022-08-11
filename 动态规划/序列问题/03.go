package main

import "fmt"

type Node struct {
	Val      byte
	Children []*Node //邻居节点
}
//返回值为答案
func process(nodeList []*Node)int{
	ret:=0
	//遍历每一个节点
	for i := 0; i < len(nodeList); i++ {
		node:=nodeList[i]
		valA:=node.Val
		//遍历node的每一个孩子   
		for j := 0; j < len(node.Children); j++ {
			childNode:=node.Children[j]
			if childNode.Val!=valA{ //abb构型，与valA不相同才行
				//遍历Child的每一个孩子
				for idx := 0; idx < len(childNode.Children); idx++ {
					sonNode:=childNode.Children[idx]
					if sonNode.Val==childNode.Val{//符合abb构型
						fmt.Printf("%c %c %c \n",valA,childNode.Val,sonNode.Val)
						
						ret++
					}
				}
			}
		}
	}
	return ret
}

func main() {
	n := 0
	fmt.Scanf("%v\n",&n)
	// fmt.Println(n)
	str:=""
	fmt.Scanf("%v\n",&str)
	// fmt.Println(str)
	//构造各个节点加入队列
	nodeList:=make([]*Node,0)
	for i := 0; i < n; i++ {
		node:=&Node{
			Val:str[i],
			Children: make([]*Node, 0),
		}
		nodeList = append(nodeList, node)
	}
	//生成树
	for j := 0; j < n-1; j++ {
		u,v:=0,0
		fmt.Scan(&u ,&v)
		u,v=u-1,v-1
		nodeList[u].Children = append(nodeList[u].Children, nodeList[v])
		nodeList[v].Children=append(nodeList[v].Children, nodeList[u])
	}
	res:=process(nodeList)
	fmt.Println(res)
}