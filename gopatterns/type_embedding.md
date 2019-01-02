# Type Embedding

In order to achieve composition, we need to take a look at type embedding in Go. Suppose we are trying
to build a car. A car is composed of many small moving parts. We want to be able to reuse these small
moving parts to build different type of cars.

## Struct Embedding

First we have a turbo charged four cylinder engine.

```go
type TurboCharged struct {
  Horsepower   int
  Torque       int
  ChargerCount int
  FanSpeed     float64
}
```

We can define couple methods on the engine.

```go
func (tc TurboCharged) ignite() {
  fmt.Printf("4 cylinder is running, %d chargers are ready\n", tc.ChargerCount)
}
```

Now let's build a car with this turbo charged engine by embedding it into a car struct.

```go
type Car struct {
  TurboCharged
}

func (c Car) Start() {
  c.ignite()
}

func main() {
  c := Car{TurboCharged{300, 350, 2, 100.5}}
  c.Start()
}
```

Output:

```text
4 cylinder is running, 2 chargers are ready
```

Notice that `ignite` is also embedded onto `Car`. I made it private so that starting an engine is
abstracted away from user.

## Interface Embedding

Now you may ask, what if I want a different engine like a naturally aspirated 6 cylinder? You can
achieve it by embedding interfaces.

```go
type Engine interface {
  ignite()
}

type NatAspirated struct {
  Horsepower int
  Torque     int
}

func (na NatAspirated) ignite() {
  fmt.Printf("6 cylinder is running, VROOOOMMMMM\n")
}
```

Change the embedded type to `Engine`

```go
type Car struct {
  Engine
}

func (c Car) Start() {
  c.ignite()
}
```

Now we can put whatever engine we like into our car.

```go
func main() {
  var c Car
  
  c = Car{NatAspirated{400, 450}}
  // OR
  c = Car{TurboCharged{300, 305, 2, 100.5}}

  c.Start()
}