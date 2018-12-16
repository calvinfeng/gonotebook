package ttt

// Player participates in a Tic Tac Toe game. It has name and mark as getters. It returns a move
// when asked.
type Player interface {
	GetMove(*board) (int, int, error)
	Mark() string
	Name() string
}
