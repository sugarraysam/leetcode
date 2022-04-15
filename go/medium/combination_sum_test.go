package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinationSum_combination_sum(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("candidates: %v", tc.candidates), func(t *testing.T) {
			t.Parallel()
			got := combinationSum(tc.candidates, tc.target)
			for _, gotCombination := range got {
				require.Contains(t, tc.expected, gotCombination)
			}
		})
	}
}
