package leetcode


func uniquePathsWithObstacles_1(obstacleGrid [][]int) int {
	x:=len(obstacleGrid)-1
	y:=len(obstacleGrid[0])-1
	return help_1(obstacleGrid,x,y)
}

//从0,0 走到x,y有多少路径
func help_1(obstacleGrid [][]int,x,y int)int{
	if obstacleGrid[x][y]==1{
		return 0
	}
    if x==0&&y==0{
        if obstacleGrid[0][0]==0{
            return 1
        }else {
			return 0
		}
    }
	if x==0{
		return help_1(obstacleGrid,x,y-1)
	}
	if y==0{
		return help_1(obstacleGrid,x-1,y)
	}
	return help_1(obstacleGrid,x-1,y)+help_1(obstacleGrid,x,y-1)
}