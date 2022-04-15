package medium

// [1,2,3]
// 2*2*2 = 2^n subsets, each subset len == n so O(n * 2^n)

func findSubsets(nums []int) [][]int {
	sol := NewSubsetSol(nums)
	sol.solve(0, []int{})
	return sol.subsets
}

type SubsetSol struct {
	nums    []int
	subsets [][]int
}

func NewSubsetSol(nums []int) *SubsetSol {
	return &SubsetSol{
		nums:    nums,
		subsets: make([][]int, 0),
	}
}

func (s *SubsetSol) solve(index int, subset []int) {
	// base case
	if index == len(s.nums) {
		s.subsets = append(s.subsets, subset)
		return
	}

	// #1) include elem
	s.solve(index+1, append(copySubset(subset), s.nums[index]))

	// #2) don't include elem
	s.solve(index+1, copySubset(subset))
}

func copySubset(in []int) []int {
	out := make([]int, len(in), len(in))
	copy(out, in)
	return out
}
