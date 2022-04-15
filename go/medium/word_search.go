package medium

func wordExistsInBoard(board [][]byte, word string) bool {
	b := NewBoard(board, word)
	return b.solve()
}

type Tile struct {
	i   int
	j   int
	val byte
}

func NewTile(i int, j int, val byte) Tile {
	return Tile{i, j, val}
}

type Board struct {
	data    [][]byte
	word    string
	rows    int
	cols    int
	visited map[Tile]bool
}

func NewBoard(data [][]byte, word string) *Board {
	return &Board{
		data:    data,
		word:    word,
		rows:    len(data),
		cols:    len(data[0]),
		visited: make(map[Tile]bool),
	}
}

func (b *Board) Neighbors(t Tile) []Tile {
	var res []Tile

	down := []int{1, 0}
	up := []int{-1, 0}
	right := []int{0, 1}
	left := []int{0, -1}

	for _, direction := range [][]int{down, up, right, left} {
		i := t.i + direction[0]
		j := t.j + direction[1]
		// inbound
		if i < b.rows && i >= 0 && j < b.cols && j >= 0 {
			neighbor := NewTile(i, j, b.data[i][j])
			res = append(res, neighbor)
		}
	}
	return res
}

// from every position (n*m) I do len(word) steps, so we have O(n * m * len(word))
func (b *Board) solve() bool {
	// not enough tiles for the word
	if b.rows*b.cols < len(b.word) {
		return false
	}

	// single char edge case
	if b.rows == 1 && b.cols == 1 {
		return b.data[0][0] == b.word[0]
	}

	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			start := NewTile(i, j, b.data[i][j])
			if b.findWord(start, 0) {
				return true
			}

			// reset visited
			b.visited = make(map[Tile]bool)
		}
	}
	return false
}

func (b *Board) findWord(curr Tile, searchIndex int) bool {
	// base case found the word
	if searchIndex == len(b.word) {
		return true
	}

	// already visited or wrong letter
	if b.visited[curr] || b.word[searchIndex] != curr.val {
		return false
	}

	// found another letter - look at neighbors recursively
	b.visited[curr] = true
	for _, neighbor := range b.Neighbors(curr) {
		if b.findWord(neighbor, searchIndex+1) {
			return true
		}
		// unvisit neighbor
		b.visited[neighbor] = false
	}
	// unvisit current
	b.visited[curr] = false
	return false
}
