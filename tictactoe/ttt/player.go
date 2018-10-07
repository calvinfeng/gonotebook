package ttt

type Player interface {
	GetMove(*Board) (int, int, error)
	Mark() string
	Name() string
}
