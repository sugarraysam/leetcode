package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test3sum_3sum(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{},
			expected: [][]int{},
		},
		{
			nums:     []int{0},
			expected: [][]int{},
		},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{-2, 0, 1, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
	}
	for _, tc := range cases {
		got := threeSum(tc.nums)
		require.Equal(t, got, tc.expected)
	}
}
