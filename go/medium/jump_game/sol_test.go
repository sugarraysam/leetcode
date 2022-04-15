package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanJump(t *testing.T) {
	cases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("nums: %v", tc.nums), func(t *testing.T) {
			t.Parallel()
			got := canJump(tc.nums)
			gotGreedy := canJumpGreedy(tc.nums)
			require.Equal(t, got, tc.expected)
			require.Equal(t, gotGreedy, tc.expected)
		})
	}
}
