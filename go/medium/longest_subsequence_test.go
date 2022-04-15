package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLCS_longest_subsequence(t *testing.T) {
	cases := []struct {
		text1    string
		text2    string
		expected int
	}{
		{
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.text1, func(t *testing.T) {
			t.Parallel()
			got := longestCommonSubsequence(tc.text1, tc.text2)
			require.Equal(t, tc.expected, got)
		})
	}
}
