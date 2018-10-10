# Factory Method
> Factory method pattern is a creational pattern that uses factory methods to deal with the problem
> of cerating objects without having to specify the exact class of the object that will be created.

For example, I define an interface Pokemon and I have multiple types of pokemon that can satisfy
this interface.
```golang
package pokemon

// Pokemon has a list of moves and attack.
type Pokemon interface {
	Moves() []string
	Attack() error
	Level() uint
}

type charmander struct {
	// List of attributes
}

func (c *charmander) Moves() []string {
	return []string{}
}

func (c *charmander) Attack() error {
	return nil
}

func (c *charmander) Level() uint {
	return c.level
}

type squirtle struct {
	// List of attributes
}

func (s *squirtle) Moves() []string {
	return []string{}
}

func (s *squirtle) Attack() error {
	return nil
}

func (s *squirtle) Level() uint {
	return s.level
}
```

I am going to use a factory method to create `Pokemon`.
```golang
// Pokemons
const (
	Charmander = "CHARMANDER"
	Squirtle   = "SQUIRTLE"
)

// NewPokemon is a factory method.
func NewPokemon(p string) Pokemon {
	switch p {
	case Charmander:
		return &charmander{}
	case Squirtle:
		return &squirtle{}
	default:
		return nil
	}
}
```

