package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	l, r, row := 0, len(matrix)-1, 0
	for l <= r {
		m := (l + r) >> 1
		if matrix[m][0] < target {
			row = m
			l = m + 1
		} else if matrix[m][0] > target {
			r = m - 1
		} else {
			return true
		}
	}
	l, r = 0, len(matrix[row])-1
	for l <= r {
		m := (l + r) >> 1
		if matrix[row][m] < target {
			l = m + 1
		} else if matrix[row][m] > target {
			r = m - 1
		} else {
			return true
		}
	}
	return false
}

func searchMatrixDeprecated(matrix [][]int, target int) bool {
	var col int
	for i := 1; i < len(matrix) && target >= matrix[i][0]; i++ {
		col = i
	}
	for i := 0; i < len(matrix[col]); i++ {
		if target == matrix[col][i] {
			return true
		}
	}
	return false
}
