# Functional Options
Golang does not provide optional arguments. Thus, we need to be a little creative here. A Porsche 
has four fields that are configurable.
```golang
type Porsche struct {
    Model  string
    Trim   string
    Color  string
    Wheels string
}
```

Let's create a function type that is basically a setter for Porsche. Along with it, we will create
four setter factory, one for each field.
```golang
type Option func(*Porsche)

func Model(val string) Option {
	return func(p *Porsche) {
		p.Model = val
	}
}

func Trim(val string) Option {
	return func(p *Porsche) {
		p.Trim = val
	}
}

func Color(val string) Option {
	return func(p *Porsche) {
		p.Color = val
	}
}

func Wheels(val string) Option {
	return func(p *Porsche) {
		p.Wheels = val
	}
}
```

The constructor will use the setters to optionally set properties for Porsche.
```golang
func NewPorsche(setters ...Option) *Porsche {
	// Default configuration
	p := &Porsche{
		Model:  "911",
		Trim:   "Base",
		Color:  "GT Silver",
		Wheels: "Carrera",
	}

	for _, setter := range setters {
		setter(p)
	}

	return p
}
```

The expected usage will look like this.
```golang
func main() {
	NewPorsche(Trim("Turbo S"), Color("Carmine Red"))
}
```