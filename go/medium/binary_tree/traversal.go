package binary_tree

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// my two lists are equal, same amount of node no matter the order
func buildTree(preorder []int, inorder []int) *TreeNode {
	// base case - no more nodes
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}

	// error case - only one list is 0 (imbalance detected)
	if len(preorder) == 0 || len(inorder) == 0 {
		log.Fatal("error imbalanced tree")
	}

	// base case - only a single node left
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	// at least 2 nodes left
	// find root position inorder
	rootVal := preorder[0]
	preorder = preorder[1:]
	rootPos := 0
	for rootPos < len(inorder) {
		if inorder[rootPos] == rootVal {
			break
		}
		rootPos++
	}

	if rootPos == len(inorder) {
		log.Fatal("could not find root inorder")
	}

	// left tree
	leftPre := preorder[:rootPos]
	leftIn := inorder[:rootPos]

	// right tree - check if rootPos == len(inorder) - 1
	rightPre := []int{}
	rightIn := []int{}
	if rootPos != len(inorder)-1 {
		rightPre = preorder[rootPos:]
		rightIn = inorder[rootPos+1:] // skip root
	}

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(leftPre, leftIn),
		Right: buildTree(rightPre, rightIn),
	}
}

func (t *TreeNode) bfs() []int {
	res := make([]int, 0)
	toVisit := []*TreeNode{t}

	for len(toVisit) > 0 {
		// pop
		curr := toVisit[0]
		toVisit = toVisit[1:]

		// case non nil
		if curr != nil {
			res = append(res, curr.Val)
			toVisit = append(toVisit, curr.Left)
			toVisit = append(toVisit, curr.Right)
		}
	}

	return res
}
