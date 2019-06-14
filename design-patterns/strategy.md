# Strategy

Strategy pattern also known as policy pattern is a behavioral software design pattern that enables selecting an algorithm at runtime.

The idea is actually very simple. It's probably something you and I have been doing the whole time.

```go
type (
    Engine interface {
        Output() float64
    }

    NaturalAspirated struct {
        // Properties
    }

    Turbocharged struct {
        // Properties
    }
)

func (na NaturalAspirated) Output() float64 {
    // Implementation...
    return power
}

func (tc Turbocharged) Output() float64 {
    // Implementation...
    return power
}
```

We have two types of engine and they both have different outputs depending on their displacement, torque curves and bunch of other details. We will have a `Car` object which takes one of these two engines and `Drive()` in runtime.

```go
type Car struct {
    Engine Engine
}

func (c *Car) Drive() {
    power := c.Engine.Output()

    // Do complex logic with engine power
}
```

