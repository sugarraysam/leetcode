package binary_tree

import "math"

// nolint
func isValidBST(root *TreeNode) bool {
	return validBSTHelper(root, math.MinInt, math.MaxInt)
}

func validBSTHelper(root *TreeNode, low int, high int) bool {
	if root == nil {
		return true
	}

	// invalid BST, (low, high), excludes low + high
	if root.Val <= low || root.Val >= high {
		return false
	}

	leftRes := validBSTHelper(root.Left, low, min(high, root.Val))
	rightRes := validBSTHelper(root.Right, max(low, root.Val), high)

	return leftRes && rightRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
