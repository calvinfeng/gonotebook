package main

func main() {
	p1 := &HumanPlayer{"Calvin", "X"}
	// p2 := &HumanPlayer{"Carmen", "O"}
	cp := &ComputerPlayer{"HAL9000", "O"}
	g := NewGame(p1, cp)
	g.Start()
}
