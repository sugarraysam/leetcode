package medium

func combinationSum(candidates []int, target int) [][]int {
	sol := NewSolCombination(candidates)
	sol.backtrack(0, target, []int{})
	return sol.combinations
}

type SolCombination struct {
	combinations [][]int
	candidates   []int
}

func NewSolCombination(candidates []int) *SolCombination {
	return &SolCombination{
		combinations: make([][]int, 0),
		candidates:   candidates,
	}
}

func (s *SolCombination) backtrack(cIndex int, target int, combination []int) {
	// found a valid combination
	if target == 0 {
		s.combinations = append(s.combinations, copySubset(combination))
		return
	}
	// deadend, we are over our target or we don't have anymore candidates
	if target < 0 || cIndex >= len(s.candidates) {
		return
	}

	// either want to include the current candidate in the decision path
	// or want to exclude it
	// # 1 - don't include current candidate anymore
	s.backtrack(cIndex+1, target, combination)

	// #2 - include current candidate
	candidate := s.candidates[cIndex]
	s.backtrack(cIndex, target-candidate, append(combination, candidate))

}
