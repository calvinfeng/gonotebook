# Tic Tac Toe

Let's build a terminal based Tic Tac Toe game in Go!

## Setup

Create a folder named `tictactoe` in your `go-academy` directory. Then create a `main.go` inside it. We will put the game logic into a package named `ttt`. For now, just make an empty `ttt` folder inside `tictactoe`.

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

`Game`: This is a struct that keeps track of the turns, the players, and the board information.

`Player`: This is an interface that has the following methods.

* `GetMove(*board) (int, int, error)`
* `Mark() string`
* `Name() string`

`HumanPlayer` and `ComputerPlayer`

`board`: This is an 3 by 3 array that keeps track of `X` and `O` marks, using `_` to render empty space.

## Quick Note

### Private vs Public

Functions or variables with upper-cased name are exported from a package. If the functions or variables have lower-cased name, then program outside of the package cannot have reference to them. I chose to make `board` a private struct because user need not to care about the implementation and usage of game board.

## Project Tic Tac Toe

{% embed url="https://www.youtube.com/watch?v=CqsfJw4HJyA&feature=youtu.be" %}

{% embed url="https://www.youtube.com/watch?v=EgpXNBqmhP8&feature=youtu.be" %}

### **Bonus**

Here are some recommendations for projects you can work on if you want more practice.

* Connect 4
* Minesweeper
* Sudoku

## Source

[GitHub](https://github.com/calvinfeng/go-academy/tree/master/tictactoe)

