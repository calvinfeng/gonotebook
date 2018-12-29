package ttt

// NewComputerPlayer is a constructor for computer player.
func NewComputerPlayer(n, m string) *ComputerPlayer {
	return &ComputerPlayer{
		name: n,
		mark: m,
	}
}

// ComputerPlayer is an AI player that makes undefeatable moves.
type ComputerPlayer struct {
	name string
	mark string
}

// GetMove returns next move.
func (cp *ComputerPlayer) GetMove(b *board) (int, int, error) {
	move := cp.minimax(b, cp.Mark(), 1)
	return move.i, move.j, nil
}

// Mark is a getter for player's mark.
func (cp *ComputerPlayer) Mark() string {
	return cp.mark
}

// Name is a getter for player's name.
func (cp *ComputerPlayer) Name() string {
	return cp.name
}

type move struct {
	value int
	i     int
	j     int
}

func (cp *ComputerPlayer) minimax(b *board, mark string, depth int) move {
	if b.winner() != "" || b.emptyCount() == 0 {
		m := move{}
		if b.winner() == cp.Mark() {
			m.value = 10 - depth
		} else {
			m.value = depth - 10
		}

		return m
	}

	moves := []move{}
	for _, pos := range b.getAvailablePos() {
		newBoard := b.copy()
		i, j := pos[0], pos[1]
		newBoard[i][j] = mark

		var opponent string
		if mark == "X" {
			opponent = "O"
		} else {
			opponent = "X"
		}

		m := cp.minimax(newBoard, opponent, depth+1)
		m.i = i
		m.j = j

		moves = append(moves, m)
	}

	// maximize move value
	if mark == cp.Mark() {
		max := moves[0]
		for _, m := range moves {
			if max.value < m.value {
				max = m
			}
		}

		return max
	}

	// minimize move value
	min := moves[0]
	for _, m := range moves {
		if min.value > m.value {
			min = m
		}
	}

	return min
}
