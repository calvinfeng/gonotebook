package ttt

// newBoard is a constructor for an empty board.
func newBoard() *board {
	var b board
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = "_"
		}
	}

	return &b
}

// board is a 3 by 3 grid.
type board [3][3]string

// String returns the string representation of a board.
func (b *board) String() string {
	str := "    "
	for i := range b {
		for j := range b[i] {
			str += b[i][j] + " "
		}
		str += "\n    "
	}
	return str
}

func (b *board) isWon() bool {
	if b.winner() != "" {
		return true
	}
	return false
}

// Winner returns the winner of the board.
func (b *board) winner() string {
	xStreak := [3]string{"X", "X", "X"}
	oStreak := [3]string{"O", "O", "O"}

	rows := b.rows()
	for i := range rows {
		if rows[i] == xStreak {
			return "X"
		}
		if rows[i] == oStreak {
			return "O"
		}
	}

	columns := b.columns()
	for i := range columns {
		if columns[i] == xStreak {
			return "X"
		}
		if columns[i] == oStreak {
			return "O"
		}
	}

	diagonals := b.diagonals()
	for i := range diagonals {
		if diagonals[i] == xStreak {
			return "X"
		}
		if diagonals[i] == oStreak {
			return "O"
		}
	}

	return ""
}

func (b *board) emptyCount() int {
	count := 0
	for i := range b {
		for j := range b[i] {
			if b[i][j] == "_" {
				count++
			}
		}
	}

	return count
}

func (b *board) rows() [3][3]string {
	return *b
}

func (b *board) columns() [3][3]string {
	columns := [3][3]string{}
	for i := range b {
		for j := range b[i] {
			columns[j][i] = b[i][j]
		}
	}
	return columns
}

func (b *board) diagonals() [2][3]string {
	diagonals := [2][3]string{}
	for i := range b {
		diagonals[0][i] = b[i][i]
	}

	for i := range b {
		diagonals[1][i] = b[i][2-i]
	}

	return diagonals
}

// copy creates a copy of the original board.
func (b board) copy() *board {
	return &b
}

// getAvailablePos returns all empty spots of the board. This is needed for bonus phase: Minimax
// algorithm
func (b *board) getAvailablePos() [][2]int {
	availPos := [][2]int{}
	for i := range b {
		for j := range b[i] {
			if b[i][j] == "_" {
				availPos = append(availPos, [2]int{i, j})
			}
		}
	}

	return availPos
}
