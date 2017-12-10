package main

func main() {
	p1 := &HumanPlayer{"Calvin", "X"}
	p2 := &HumanPlayer{"Carmen", "O"}
	g := NewGame(p1, p2)
	g.Start()
}
