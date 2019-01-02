# Tic Tac Toe

Let's build a terminal based Tic Tac Toe game in Go!

## Setup

Create a folder named `tictactoe` in your `go-academy` directory. Then create a `main.go` inside it.
We will put the game logic into a package named `ttt`. For now, just make an empty `ttt` folder
inside `tictactoe`.

```text
WORKSPACE/
    src/
        go-academy/
            tictactoe/
                main.go
                ttt/
                    ...
```

## Object Oriented Programming in Go

Let's define what are entities that we need for a terminal based Tic Tac Toe game.

1. `Game`: This is a struct that keeps track of the turns, the players, and the board information.
2. `Player`: This is an interface that has the following methods.
    * `GetMove(*board) (int, int, error)`
    * `Mark() string`
    * `Name() string`
3. `HumanPlayer` and `ComputerPlayer`
4. `board`: This is an 3x3 array that keeps track of `X` and `O` marks, using `_` to render empty space.

## Quick Note

### Private vs Public

Functions or variables with upper-cased name are exported from a package. If the functions or variables
have lower-cased name, then program outside of the package cannot have reference to them. I chose to
make `board` a private struct because user need not to care about the implementation and usage of
game board.

## Tic Tac Toe in Go

* [Lesson 2 Tic Tac Toe Introduction](https://youtu.be/CqsfJw4HJyA)
* [Lesson 2 Tic Tac Toe Implementation](https://youtu.be/EgpXNBqmhP8)

## Homework

Here are some recommendations for projects you can work on if you want more practice.

* Connect 4
* Minesweeper
* Sudoku

## Source

[GitHub](https://github.com/calvinfeng/go-academy/tree/master/tictactoe)