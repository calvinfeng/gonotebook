# Tic Tac Toe
Let's build a terminal based Tic Tac Toe game in Go!

## Setup
Create a folder named `tictactoe` in your `go-academy` directory. Then create a `main.go` inside it.
We will put the game logic into a package named `ttt`. For now, just make an empty `ttt` folder
inside `tictactoe`. 
```
WORKSPACE/
    src/
        go-academy/
            tictactoe/
                main.go
                ttt/
                    ...
```

## Quick Note
### Private vs Public
Functions or variables with upper-cased name are exported from a package. If the functions or 
variables have lower-cased name, then program outside of the package cannot have reference to them.

## Object Oriented Programming in Go
Let's define what are entities that we need for a terminal based Tic Tac Toe game. 

1. `Game`: This is a struct that keeps track of the turns, the players, and the board information.
2. `board`: This is a struct that keeps track of `X` and `O` marks. 
3. `Player`: This is an interface that implements the following methods
    * `GetMove(*board) (int, int, error)`
    * `Mark() string`
    * `Name() string`
4. `HumanPlayer` and `ComputerPlayer`

### Game
The data structure should look like the following.
```golang
type Game struct {
    p1 Player 
    p2 Player
    current Player
    board *board 
    round int
}
```

It should have one exported method and bunch of other private methods for internal logic.
```
func (g *Game) Start() {
    // Implementation...
}
```

### Board
`board` is probably the least pleasant piece in this project because it contains logic to check 
columns, rows, and diagonals. We can define board as a 3 by 3 array of string.
```golang
type board [3][3]string
```

It should have the following API(s).
* `String() string`
* `winner() string`
* `getEmptyCount() int`

For bonus phase, we also need the following.
* `copy() *board`
* `getAvailablePos() [][2]int`

There is a very neat trick with doing a deep copy of an array. We can declare the function as value 
receiver, which means it takes the value of pointer to `board` and copies into the function. Thus, 
the `b` inside the function is already a copy of the original board. We just need to take its pointer
address and return it.
```
func (b board) copy() *board {
    return *b
}
```

### Player Interface
Each player should have name and mark. Mark is either `X` or `O`. 
```golang
type Player interface {
    GetMove(*board) (int, int, error)
    Mark() string
    Name() string
}
```

### Human Player
The main responsibility of a `HumanPlayer` struct is to obtain user input from terminal. And 
of course, it should implement the `Player` interface.
```golang 
type HumanPlayer struct {
	name string
	mark string
}
```

For example, we can use `fmt.Scanf` to get user inputs.
```golang
// GetMove returns next move.
func (p *HumanPlayer) GetMove(b *board) (int, int, error) {
	fmt.Print("Enter position: ")
	var i, j int
	if n, err := fmt.Scanf("%d %d", &i, &j); err != nil || n != 2 {
		return 0, 0, err
	}

	fmt.Println("Your input:", i, j)
	return i, j, nil
}
```

## Tic Tac Toe in Go
* [Tic Tac Toe in Go](https://www.youtube.com/channel/UCoKwJSadNdeJkpfBpI-f5Ow)

## Bonus
Here are some recommendations for projects you can work on if you want more practice.-
* Connect 4
* Minesweeper
* Sudoku

## Source
[GitHub](https://github.com/calvinfeng/go-academy/tree/master/tictactoe)



