package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestAbsPath_longest_bas_path(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			input:    "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			expected: 20,
		},
		{
			input:    "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			expected: 32,
		},
		{
			input:    "a",
			expected: 0,
		},
		{
			input:    "dir\nfile.txt",
			expected: 8,
		},
	}
	for _, tc := range cases {
		got := lengthLongestPath(tc.input)
		require.Equal(t, got, tc.expected)
	}
}
