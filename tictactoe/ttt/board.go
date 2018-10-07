package ttt

// NewBoard is a constructor for an empty board.
func NewBoard() *Board {
	newBoard := &Board{}
	for i := range newBoard.Grid {
		for j := range newBoard.Grid[i] {
			newBoard.Grid[i][j] = "_"
		}
	}
	return newBoard
}

// Board is a 3 by 3 grid.
type Board struct {
	Grid [3][3]string
}

// Copy creates a copy of the original board.
func (b Board) Copy() *Board {
	return &b
}

// PlaceMark places a mark on a given spot on the grid.
func (b *Board) PlaceMark(i, j int, mark string) {
	b.Grid[i][j] = mark
}

// IsOver checks if a game is over.
func (b *Board) IsOver() bool {
	return b.isWon() || b.isTied()
}

// String returns the string representation of a board.
func (b Board) String() string {
	boardStr := "    "
	for i := range b.Grid {
		for j := range b.Grid[i] {
			boardStr += b.Grid[i][j] + " "
		}
		boardStr += "\n    "
	}
	return boardStr
}

func (b *Board) isWon() bool {
	if b.Winner() != "" {
		return true
	}
	return false
}

// Winner returns the winner of the board.
func (b *Board) Winner() string {
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

func (b *Board) isTied() bool {
	spaceCount := 0
	for i := range b.Grid {
		for j := range b.Grid[i] {
			if b.Grid[i][j] == "_" {
				spaceCount += 1
			}
		}
	}

	if spaceCount == 0 && b.Winner() == "" {
		return true
	}
	return false
}

func (b *Board) rows() [3][3]string {
	return b.Grid
}

func (b *Board) columns() [3][3]string {
	columns := [3][3]string{}
	for i := range b.Grid {
		for j := range b.Grid[i] {
			columns[j][i] = b.Grid[i][j]
		}
	}
	return columns
}

func (b *Board) diagonals() [2][3]string {
	diagonals := [2][3]string{}
	for i := range b.Grid {
		diagonals[0][i] = b.Grid[i][i]
	}

	for i := range b.Grid {
		diagonals[1][i] = b.Grid[i][2-i]
	}

	return diagonals
}

// GetAvailablePos returns all empty spots of the board. This is needed for bonus phase: Minimax
// algorithm
func (b *Board) GetAvailablePos() [][2]int {
	availPos := [][2]int{}
	for i := range b.Grid {
		for j := range b.Grid[i] {
			if b.Grid[i][j] == "_" {
				availPos = append(availPos, [2]int{i, j})
			}
		}
	}

	return availPos
}
