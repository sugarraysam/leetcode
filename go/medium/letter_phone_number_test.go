package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLetterCombinations_letter_phone_number(t *testing.T) {
	cases := []struct {
		digits   string
		expected []string
	}{
		{
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits:   "",
			expected: []string{},
		},
		{
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.digits, func(t *testing.T) {
			t.Parallel()
			got := letterCombinations(tc.digits)
			require.Equal(t, got, tc.expected)
		})
	}
}
