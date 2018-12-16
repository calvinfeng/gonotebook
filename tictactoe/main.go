package main

import "go-academy/tictactoe/ttt"

func main() {
	hp := ttt.NewHumanPlayer("Calvin", "X")
	cp := ttt.NewComputerPlayer("HAL9000", "O")
	g := ttt.NewGame(hp, cp)
	g.Start()
}
