package medium

import (
	"path"
	"strings"
)

// n files (>= 1)
// root dir
// m >= 0, number of root directories

// res, list of groups (slice of slice), of duplicate file paths
// each group is list of filepaths with same content

type File struct {
	Path    string
	Content string
}

// map[string][]string - where key is file content, []string is all paths that
// share this file content
// iterate through map, if len([]string) > 1, than this is a duplicate group

// case m == 0 (no root), returns empty [][]string (OK)
func findDuplicate(paths []string) [][]string {
	var res [][]string
	contentToPaths := make(map[string][]string)

	for _, path := range paths {
		files := parseFiles(path)
		for _, f := range files {
			contentToPaths[f.Content] = append(contentToPaths[f.Content], f.Path)
		}
	}

	// iterate and find duplicates
	for _, paths := range contentToPaths {
		if len(paths) > 1 {
			res = append(res, paths)
		}
	}
	return res
}

// assume m >= 1
func parseFiles(path string) []File {
	var res []File
	parts := strings.Split(path, " ")

	// n >= 1, we always have a file
	root := parts[0]
	for i := 1; i < len(parts); i++ {
		f := newFile(root, parts[i])
		res = append(res, f)
	}

	return res
}

// rawFile == "1.txt(content)"
func newFile(root string, rawFile string) File {
	parts := strings.Split(rawFile, "(")

	filePath := parts[0]
	content := strings.TrimRight(parts[1], ")")

	return File{
		Path:    path.Join(root, filePath),
		Content: content,
	}
}
