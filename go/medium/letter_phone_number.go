package medium

import (
	"log"
	"strconv"
	"strings"
)

func letterCombinations(digits string) []string {
	// empty case
	if len(digits) == 0 {
		return []string{}
	}

	digitToLetters := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	var letters [][]string
	for i := 0; i < len(digits); i++ {
		d, err := strconv.Atoi(string(digits[i]))
		if err != nil {
			log.Fatal(err)
		}
		letters = append(letters, digitToLetters[d])
	}

	sol := NewSolution()
	sol.solve(letters, []string{})
	return sol.combinations
}

type Solution struct {
	combinations []string
}

func NewSolution() *Solution {
	return &Solution{
		combinations: make([]string, 0),
	}
}

func (s *Solution) solve(letters [][]string, combination []string) {
	// we have constructed a valid combination
	if len(combination) == len(letters) {
		s.combinations = append(s.combinations, strings.Join(combination, ""))
		return
	}

	depth := len(combination)
	for i := 0; i < len(letters[depth]); i++ {
		combination = append(combination, letters[depth][i])
		s.solve(letters, combination)

		// remove last element
		combination = combination[:len(combination)-1]
	}
}
