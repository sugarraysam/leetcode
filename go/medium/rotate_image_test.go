package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateImage_rotate_image(t *testing.T) {
	t.Parallel()
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}

	type testFunc func([][]int) [][]int

	for _, fn := range []testFunc{RotateImageClockWise, RotateImageCounterClockWise} {
		matrixCopy := make([][]int, len(matrix), len(matrix))
		for i := range matrix {
			matrixCopy[i] = make([]int, len(matrix[i]), len(matrix[i]))
			copy(matrixCopy[i], matrix[i])
		}
		require.Equal(t, fn(matrixCopy), expected)
	}
}
