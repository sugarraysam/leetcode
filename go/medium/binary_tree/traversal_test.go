package binary_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildTree_traversal(t *testing.T) {
	cases := []struct {
		preorder []int
		inorder  []int
		expected []int
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: []int{3, 9, 20, 15, 7}, // bfs traversal
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("pre: %v", tc.preorder), func(t *testing.T) {
			t.Parallel()
			tree := buildTree(tc.preorder, tc.inorder)
			require.Equal(t, tree.bfs(), tc.expected)
		})
	}
}
