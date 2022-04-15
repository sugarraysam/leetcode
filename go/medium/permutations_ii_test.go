package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermutationsII_permutations_ii(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{1, 1, 2},
			expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("nums %v", tc.nums), func(t *testing.T) {
			t.Parallel()
			got := permuteUnique(tc.nums)
			fmt.Println(got)
			for _, gotPerm := range got {
				require.Contains(t, tc.expected, gotPerm)
			}
		})
	}
}
