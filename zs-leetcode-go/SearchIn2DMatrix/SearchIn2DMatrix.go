package SearchIn2DMatrix

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	x, y := 0, cols-1
	for x < rows && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
