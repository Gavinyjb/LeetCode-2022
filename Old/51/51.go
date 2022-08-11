package leetcode_51

//package main

func solveNQueens(n int) [][]string {
	res:=make([][]int,0)
	track:=make([]int,0)

	var backtrack func(n int,track []int)
	backtrack=func(n int, track []int) {
		//触发结束条件
		if len(track)==n{
			temp:=make([]int,n)
			copy(temp,track)
			res=append(res,temp)
		}
		//for 选择 in 选择列表:
		for i := 0; i < n; i++ {
			//排除不合法的选择
			if !isValid(track,n,i){
				//该位置不可放置
				continue
			}else {
				//做选择
				track=append(track,i)
				//进入下一层决策树
				backtrack(n,track)
				//取消选择
				track=track[:len(track)-1]
			}
		}
	}
	backtrack(n,track)
	return printRes(res,n)
}
//isValid 使用len(track)-1 代替row 行
func isValid(track []int,n int,col int)bool{
	//生成棋盘
	board:=make([][]int,0)
	for _, v := range track {
		row:=make([]int,n)
		for j := 0; j < n; j++ {
			if j==v{
				row[j]=1
			}else {
				row[j]=0
			}
		}
		board=append(board,row)
	}
	//检查是否列有冲突
	for i := 0; i < len(track); i++ {
		if track[i]==col{
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i,j:= len(track)-1,col+1 ; i >=0 && j<n ; {
		if board[i][j]==1{
			return false
		}
		i--
		j++
	}
	// 检查左上方是否有皇后互相冲突
	for i,j:= len(track)-1,col-1;i>=0&&j>=0;{
		if board[i][j]==1{
			return false
		}
		i--
		j--
	}
	return true
}

//画棋盘
func printRes(res [][]int,n int) [][]string {
	anss:=make([][]string,0)


	for _, re := range res {
		ans:=make([]string,0)
		for _, v := range re {
			row:=""
			for i := 0; i < n; i++ {
				if i==v{
					row+="Q"
				}else {
					row+="."
				}
			}
			ans = append(ans, row)
		}
		anss = append(anss, ans)
	}
	return anss
}

//func main()  {
//	res:=solveNQueens(4)
//	fmt.Println(res)
//
//}
