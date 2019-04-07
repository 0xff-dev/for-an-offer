package SetMatrixZeroes

import "fmt"

func setZeroes(matrix [][]int) {
	isCols := false
	i, j := 0, 0
	rows, cols := len(matrix), len(matrix[0])
	for ; i < rows; i++ {
		if matrix[i][0] == 0 {
			isCols = true
		}
		for j = 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i = 1; i < rows; i++ {
		for j = 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for j = 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}
	if isCols {
		for i = 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
	fmt.Println(matrix)
}
