package medium

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransposeFile_transpose_file(t *testing.T) {
	rootDir := ".."
	testFilePath := "testdata/medium/transpose_file.txt"

	absPath, err := filepath.Abs(path.Join(rootDir, testFilePath))
	require.Nil(t, err)

	expected := "name alice ryan\nage 21 30\n"

	got := transposeFile(absPath)
	require.Equal(t, got, expected)
}
