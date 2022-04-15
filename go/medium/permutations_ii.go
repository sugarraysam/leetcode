package medium

func permuteUnique(nums []int) [][]int {
	sol := NewPermSol(nums)
	sol.solve([]int{})
	return sol.perms
}

type PermSol struct {
	nums    []int
	perms   [][]int
	numsMap map[int]int
}

func NewPermSol(nums []int) *PermSol {
	numsMap := make(map[int]int)
	for _, n := range nums {
		numsMap[n]++
	}

	return &PermSol{
		nums:    nums,
		perms:   make([][]int, 0),
		numsMap: numsMap,
	}
}

// recursive backtrack algorithm get all unique perms
func (p *PermSol) solve(perm []int) {
	if len(perm) == len(p.nums) {
		p.perms = append(p.perms, copySubset(perm))
		return
	}

	for candidate, count := range p.numsMap {
		if count == 0 {
			continue
		}
		perm = append(perm, candidate)
		p.numsMap[candidate]--
		p.solve(perm)

		// backtrack
		p.numsMap[candidate]++
		perm = perm[:len(perm)-1]
	}
}
