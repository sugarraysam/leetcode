package medium

import (
	"path"
	"strings"
)

// every file and directory has a unique absolute path

// dir name: letters, digist and SPACES
// file name: name.extension, letters digits and or SPACES

// parse input and get all files in abs path
// return longest length of longest abs path

// there might be no file in the system, return 0

// 1. parse input and build a tree (not a binary tree)
// 2. traverse tree find all files and get their abs path
// 3. return longest abs path or 0
func lengthLongestPath(input string) int {
	var allFiles []string

	lines := strings.Split(input, "\n")

	// empty case
	if len(lines) == 0 {
		if isFile(input) {
			return len(input)
		}
		return 0
	}

	// there can be more than one root
	// there can also be files at the root level
	i := 0
	for i < len(lines) {
		node := NewNode(lines[i], nil)
		i++

		if node.IsFile && node.Depth == 0 {
			allFiles = append(allFiles, node.Name)
			continue
		} else {
			// found a root directory
			createTreeRecursive(node, lines[i:])
			allFiles = append(allFiles, node.BFSFindAllFiles()...)
		}
	}

	// find largest file
	maxLen := 0
	for _, file := range allFiles {
		maxLen = max(maxLen, len(file))
	}
	return maxLen
}

// algorithm, when I find a dir, call recursively, otherwise, simply append child
// when depth increases, go deeper in recursion
// when depth goes back to parent depth, exit recursion
func createTreeRecursive(parent *Node, lines []string) {
	i := 0
	for i < len(lines) {
		node := NewNode(lines[i], parent)
		i++

		// depth decreases, exit recursion
		if node.Depth <= parent.Depth {
			return
		}

		// only interested in entries 1 level above
		if node.Depth-parent.Depth == 1 {
			parent.AppendChild(node)

			// deeper in recursion if its a dir
			if !node.IsFile {
				createTreeRecursive(node, lines[i:])
			}
		}
	}
}

type Node struct {
	Name   string
	Depth  int
	IsFile bool
	Childs []*Node
	Parent *Node
}

func NewNode(rawEntry string, parent *Node) *Node {
	name := strings.TrimLeft(rawEntry, "\t")
	return &Node{
		Name:   name,
		Depth:  strings.Count(rawEntry, "\t"),
		IsFile: isFile(name),
		Childs: make([]*Node, 0),
		Parent: parent,
	}
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}

// rawName: file.ext OR dir
func (n *Node) AppendChild(child *Node) {
	n.Childs = append(n.Childs, child)
}

func (n *Node) AbsPath() string {
	if n == nil {
		return ""
	}

	// walk up the tree to find all parents
	res := n.Name
	curr := n
	for curr.Parent != nil {
		curr = curr.Parent
		res = path.Join(curr.Name, res)
	}
	return res
}

func (n *Node) BFSFindAllFiles() []string {
	var res []string

	// empty case
	if n == nil {
		return res
	}

	toVisit := []*Node{n}
	for len(toVisit) > 0 {
		// visit a node
		curr := toVisit[0]
		toVisit = toVisit[1:]

		if curr.IsFile {
			res = append(res, curr.AbsPath())
		}
		toVisit = append(toVisit, curr.Childs...)
	}

	return res
}

func isFile(rawName string) bool {
	return strings.Contains(rawName, ".")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
