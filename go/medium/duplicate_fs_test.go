package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindDuplicate_duplicate_fs(t *testing.T) {
	cases := []struct {
		paths    []string
		expected [][]string
	}{
		{
			paths:    []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			expected: [][]string{{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}},
		},

		{
			paths:    []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"},
			expected: [][]string{{"root/a/2.txt", "root/c/d/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}},
		},
	}
	for _, tc := range cases {
		got := findDuplicate(tc.paths)
		for _, duplicate := range got {
			require.Contains(t, tc.expected, duplicate)
		}
	}
}
