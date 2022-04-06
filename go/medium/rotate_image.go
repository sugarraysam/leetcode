package medium

// rotate a n x n matrix in-place
func RotateImageClockWise(matrix [][]int) [][]int {
	left := 0
	right := len(matrix) - 1

	for left < right {
		for i := 0; i < right-left; i++ {
			top := left
			bottom := right

			topLeft := matrix[top][left+i]

			matrix[top][left+i] = matrix[bottom-i][left]
			matrix[bottom-i][left] = matrix[bottom][right-i]
			matrix[bottom][right-i] = matrix[top+i][right]
			matrix[top+i][right] = topLeft
		}
		right--
		left++
	}

	return matrix
}

func RotateImageCounterClockWise(matrix [][]int) [][]int {
	left := 0
	right := len(matrix) - 1

	for left < right {
		for i := 0; i < right-left; i++ {
			top := left
			bottom := right

			// tmp all values
			topLeft := matrix[top][left+i]
			topRight := matrix[top+i][right]
			bottomRight := matrix[bottom][right-i]
			bottomLeft := matrix[bottom-i][left]

			// overwrite
			matrix[top][left+i] = bottomLeft
			matrix[top+i][right] = topLeft
			matrix[bottom][right-i] = topRight
			matrix[bottom-i][left] = bottomRight

		}
		left++
		right--
	}

	return matrix
}
