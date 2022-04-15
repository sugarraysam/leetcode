package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordSearch_word_search(t *testing.T) {
	cases := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCCED",
			expected: true,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "SEE",
			expected: true,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCB",
			expected: false,
		},
		{
			board:    [][]byte{{'A'}},
			word:     "A",
			expected: true,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCESEEEFS",
			expected: true,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.word, func(t *testing.T) {
			t.Parallel()
			got := wordExistsInBoard(tc.board, tc.word)
			require.Equal(t, got, tc.expected)
		})
	}
}
