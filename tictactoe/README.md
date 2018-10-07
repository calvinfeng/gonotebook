# Tic Tac Toe
[Source code is here](https://github.com/calvinfeng/go-academy/tree/master/tictactoe)

## Object Oriented Programming in Go
Let's define what are entities that we need for a terminal based Tic Tac Toe game. 

1. `Game`: This is a struct that keeps track of the turns, the players, and the board information.
2. `Board`: This is a struct that keeps track of `X` and `O` marks. 
3. `Player`: This is an interface that implements the following methods
    * `GetMove(*Board) (int, int, error)`
    * `Mark() string`
    * `Name() string`
4. `HumanPlayer` and `ComputerPlayer`

First step, let's create a folder called `ttt` which is short for tic tac toe. Inside this folder we
will create a `ttt` package.

### Game

### Board

### Player Interface

### Human Player


## Bonus
Create a `ComputerPlayer` that implements the `Player` interface and make it undefeatable with 
*Minimax* algorithm. Here are some other recommendations for projects you can work on

* Connect 4
* Minesweeper
* Sudoku

## (Optional) Video 02: Tic Tac Toe in Go
* [Tic Tac Toe in Go Part 1](https://youtu.be/644HhokVkbI)
* [Tic Tac Toe in Go Part 2](https://youtu.be/eL6ruTgOQG0)
* [Tic Tac Toe in Go Part 3](https://youtu.be/rdSgqye50Qw)

