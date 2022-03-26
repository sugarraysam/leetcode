package medium

import (
	"bufio"
	"os"
	"strings"
)

// each row has same number of columns (might be more than 1 colum
// at least 1 column per row
// each field separated by ' '

// row := []string
// rows := [][]string{}
// rowsToColumns(rows)

// assume we have at least one row
func transposeFile(path string) string {
	rows := parseFileToRows(path)
	transposed := make([][]string, len(rows[0]))

	// Example of rows
	// [][]string{
	//   {"a", "b"},
	//	 {"c", "d"},
	//   {"e", "f"},
	// }

	// i -> 0, 1
	for i := 0; i < len(rows[0]); i++ {
		// j -> 0, 1, 2
		for j := 0; j < len(rows); j++ {
			// [][]string{
			//   {"a", "c", "e"},
			//	 {"b", "d", "f"},
			// }
			transposed[i] = append(transposed[i], rows[j][i])
		}
	}

	res := strings.Builder{}
	for _, row := range transposed {
		res.WriteString(strings.Join(row, " "))
		res.WriteString("\n")
	}
	return res.String()
}

func parseFileToRows(path string) [][]string {
	var res [][]string

	file, err := os.Open(path)
	if err != nil {
		panic("failed to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// next token is '\n'
		// all slices in res have the same length
		line := scanner.Text()
		res = append(res, strings.Split(line, " "))
	}

	if err := scanner.Err(); err != nil {
		panic("failed to read file")
	}

	return res
}
