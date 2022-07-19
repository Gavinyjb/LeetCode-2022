package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	//思路：将二维数组转为一维数组  进行二分搜索
	row := len(matrix)
	col := len(matrix[0])
	start, end := 0, row*col-1
	for start+1 < end {
		mid := start + (end-start)>>1
		//获取二维数组对应值
		val := matrix[mid/col][mid%col]
		if val > target {
			end = mid
		} else if val < target {
			start = mid
		} else {
			return true
		}
	}
	if matrix[start/col][start%col] == target {
		return true
	}
	if matrix[end/col][end%col] == target {
		return true
	}
	return false
}
