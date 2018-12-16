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
	i, j := move["i"], move["j"]
	return i, j, nil
}

// Mark is a getter for player's mark.
func (cp *ComputerPlayer) Mark() string {
	return cp.mark
}

// Name is a getter for player's name.
func (cp *ComputerPlayer) Name() string {
	return cp.name
}

func (cp *ComputerPlayer) minimax(b *board, mark string, depth int) map[string]int {
	if b.isOver() {
		var score map[string]int
		score = make(map[string]int)
		if b.winner() == cp.Mark() {
			score["value"] = 10 - depth
		} else {
			score["value"] = depth - 10
		}
		return score
	}

	scores := []map[string]int{}
	for _, pos := range b.getAvailablePos() {
		newBoard := b.copy()
		i, j := pos[0], pos[1]
		newBoard[i][j] = mark

		var score map[string]int
		if mark == "X" {
			score = cp.minimax(newBoard, "O", depth+1)
		} else {
			score = cp.minimax(newBoard, "X", depth+1)
		}
		score["i"] = i
		score["j"] = j
		scores = append(scores, score)
	}

	// max
	if mark == cp.Mark() {
		maxScore := scores[0]
		for _, s := range scores {
			if maxScore["value"] < s["value"] {
				maxScore = s
			}
		}
		return maxScore
	}

	// min
	minScore := scores[0]
	for _, s := range scores {
		if minScore["value"] > s["value"] {
			minScore = s
		}
	}

	return minScore
}
