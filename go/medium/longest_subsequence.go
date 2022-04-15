package medium

func longestCommonSubsequence(text1 string, text2 string) int {
	sol := NewSolveLCS(text1, text2)
	return sol.solve(0, 0)
}

type SolveLCS struct {
	text1 string
	text2 string
	cache map[lcsKey]int
}

type lcsKey struct {
	i int
	j int
}

func NewLcsKey(i, j int) lcsKey {
	return lcsKey{i, j}
}

func NewSolveLCS(text1, text2 string) *SolveLCS {
	return &SolveLCS{
		text1: text1,
		text2: text2,
		cache: make(map[lcsKey]int),
	}
}

func (s *SolveLCS) solve(i, j int) int {
	// base case - index out of bounds
	if i >= len(s.text1) || j >= len(s.text2) {
		return 0
	}

	// base case - entry is cached
	k := NewLcsKey(i, j)
	if v, ok := s.cache[k]; ok {
		return v
	}

	// case 1: character match, advance both ptrs, solve subproblem
	if s.text1[i] == s.text2[j] {
		res := 1 + s.solve(i+1, j+1)
		s.cache[k] = res
		return res
	}

	// case 2: character don't match - return max of both paths
	res := max(s.solve(i, j+1), s.solve(i+1, j))
	s.cache[k] = res
	return res
}
