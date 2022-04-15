package course_schedule_ii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindOrder(t *testing.T) {
	cases := []struct {
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {2, 1}, {3, 1}, {3, 2}},
			expected:      []int{0, 1, 2, 3},
		},
		{
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      []int{0},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}, {0, 1}},
			expected:      []int{},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{0, 1}, {3, 1}, {1, 3}, {3, 2}},
			expected:      []int{},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("n=%v", tc.numCourses), func(t *testing.T) {
			t.Parallel()
			got := findOrder(tc.numCourses, tc.prerequisites)
			require.Equal(t, tc.expected, got)
		})
	}
}
