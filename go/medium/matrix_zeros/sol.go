package matrix_zeros

// row 0 as column work array
// col 1 as row work array + firstRow constant
func setZeroes(matrix [][]int) {
	m := NewMatrix(matrix)
	m.solve()
}

type Matrix struct {
	firstRow int
	data     [][]int
	rows     int
	cols     int
}

func NewMatrix(data [][]int) *Matrix {
	return &Matrix{
		firstRow: 1, // do nothing by default
		data:     data,
		rows:     len(data),
		cols:     len(data[0]),
	}
}

func (m *Matrix) ZeroRow(r int) {
	for c := 0; c < m.cols; c++ {
		m.data[r][c] = 0
	}
}

func (m *Matrix) ZeroColum(c int) {
	for r := 0; r < m.rows; r++ {
		m.data[r][c] = 0
	}
}

func (m *Matrix) solve() {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			// case 1: just continue
			if m.data[r][c] != 0 {
				continue
			}

			// case 0: set associated row and col to be zeroed
			// col work == first row
			m.data[0][c] = 0

			// row work == first col + m.firstRow
			if r == 0 {
				m.firstRow = 0
			} else {
				m.data[r][0] = 0
			}
		}
	}
	// zero rows (1-n)
	for r := 1; r < m.rows; r++ {
		if m.data[r][0] == 0 {
			m.ZeroRow(r)
		}
	}

	// zero columns (1-n)
	for c := 1; c < m.cols; c++ {
		if m.data[0][c] == 0 {
			m.ZeroColum(c)
		}
	}
	// zero col 0
	if m.data[0][0] == 0 {
		m.ZeroColum(0)
	}
	// zero row 0
	if m.firstRow == 0 {
		m.ZeroRow(0)
	}
}
