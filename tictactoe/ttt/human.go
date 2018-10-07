package ttt

import "fmt"

// NewHumanPlayer is a constructor for human player.
func NewHumanPlayer(n string, m string) *HumanPlayer {
	return &HumanPlayer{
		name: n,
		mark: m,
	}
}

// HumanPlayer asks user to input to make next move.
type HumanPlayer struct {
	name string
	mark string
}

// GetMove returns next move.
func (p *HumanPlayer) GetMove(b *Board) (int, int, error) {
	fmt.Print("Enter position: ")
	var i, j int
	if n, err := fmt.Scanf("%d %d", &i, &j); err != nil || n != 2 {
		return 0, 0, err
	}

	fmt.Println("Your input:", i, j)
	return i, j, nil
}

// Mark is a getter for player's mark.
func (p *HumanPlayer) Mark() string {
	return p.mark
}

// Name is a getter for player's name.
func (p *HumanPlayer) Name() string {
	return p.name
}
