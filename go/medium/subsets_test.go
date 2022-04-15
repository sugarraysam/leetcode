package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindSubsets_find_subsets(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("findSubsets %v", tc.nums), func(t *testing.T) {
			t.Parallel()
			got := findSubsets(tc.nums)
			for _, gotSubset := range got {
				require.Contains(t, tc.expected, gotSubset)
			}
		})
	}
}
