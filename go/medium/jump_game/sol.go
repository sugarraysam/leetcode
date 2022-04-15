package medium

func canJump(nums []int) bool {
	// base case 1 value
	if len(nums) == 1 {
		return true
	}

	// len(nums) >= 2
	dp := make(map[int]bool)
	dp[len(nums)-1] = true

	// walking back
	for pos := len(nums) - 2; pos >= 0; pos-- {
		jumpLen := nums[pos]
		canReach := false

		for j := 1; j <= jumpLen; j++ {
			// out of bounds
			if pos+j >= len(nums) {
				return false
			}

			if dp[pos+j] == true {
				canReach = true
				break
			}
		}
		dp[pos] = canReach
	}

	return dp[0]
}

// greedy by moving goal post as close as possible starting from the end
// this works because you can reach all intermediary positions as well (1,2,..n)
func canJumpGreedy(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	goal := len(nums) - 1
	for pos := len(nums) - 2; pos >= 0; pos-- {
		if pos+nums[pos] >= goal {
			goal = pos
		}
	}
	return goal == 0
}
