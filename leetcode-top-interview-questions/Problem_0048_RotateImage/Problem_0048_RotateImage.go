package Problem_0048_RotateImage

func rotate(matrix [][]int) {
	upRow, leftCol, downRow, rightCol := 0, 0, len(matrix)-1, len(matrix[0])-1
	for upRow < downRow {
		processEdge(matrix, upRow, leftCol, downRow, rightCol)
		upRow++
		leftCol++
		downRow--
		rightCol--
	}

}

func processEdge(nums [][]int, upRow, leftCol, downRow, rightCol int) {
	times := downRow - upRow
	temp := 0
	for i := 0; i < times; i++ {
		temp = nums[upRow][leftCol+i]
		nums[upRow][leftCol+i] = nums[downRow-i][leftCol]
		nums[downRow-i][leftCol] = nums[downRow][rightCol-i]
		nums[downRow][rightCol-i] = nums[upRow+i][rightCol]
		nums[upRow+i][rightCol] = temp
	}
}
