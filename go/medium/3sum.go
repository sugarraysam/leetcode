package medium

import "sort"

// unique triplets that add to 0
// no duplicate triplets
// (a,b,c) where a+b+c == 0

// strategy:
// - sort array O(n logn)
// - fix the a value (must be negative and unvisited)
// - solve for 2sum to find b and c

type sol struct {
	a int
	b int
	c int
}

func NewSol(a int, bc []int) sol {
	return sol{a: a, b: bc[0], c: bc[1]}
}

func (s sol) ToArr() []int {
	return []int{s.a, s.b, s.c}
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	// sort array in place
	sort.Ints(nums)

	// keep track of solutions
	seen := make(map[sol]bool)

	for aPtr := 0; aPtr < len(nums)-1; aPtr++ {
		// nums is sorted, so we need at least one negative value in our
		// triplets to add to 0
		a := nums[aPtr]
		if a > 0 {
			return res
		}

		bcs := twoSum(-a, nums[aPtr+1:])
		for _, bc := range bcs {
			sol := NewSol(a, bc)
			if _, ok := seen[sol]; !ok {
				// new triplet
				seen[sol] = true
				res = append(res, sol.ToArr())
			}
		}
	}

	return res
}

// can have multiple solutions
func twoSum(target int, nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{}
	}

	// len(nums) >= 2
	res := make([][]int, 0)
	leftPtr := 0
	rightPtr := len(nums) - 1

	for leftPtr < rightPtr {
		leftVal := nums[leftPtr]
		rightVal := nums[rightPtr]
		sum := leftVal + rightVal

		if sum == target {
			res = append(res, []int{leftVal, rightVal})
			leftPtr++
			rightPtr--
		} else if sum < target {
			// want to increase sum
			// we know nums is a sorted array, we increase by moving the leftPtr
			leftPtr++
		} else {
			// sum > target
			// want to decrease the sum
			// nums sorted so we decrease by moving rightPtr
			rightPtr--
		}
	}
	// no solution
	return res
}
