package matrix_zeros

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrixZeroes(t *testing.T) {
	cases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix:   [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			matrix:   [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	count := 0
	for _, tc := range cases {
		tc := tc
		count++
		t.Run(fmt.Sprintf("test #%v", count), func(t *testing.T) {
			t.Parallel()
			setZeroes(tc.matrix)
			require.Equal(t, tc.matrix, tc.expected)
		})
	}
}
