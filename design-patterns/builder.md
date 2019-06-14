# Builder

We want to build some cars here. Let's see we can use builder pattern build some cars.

> Builder pattern separates the construction of a complex object from its representation so that the same construction process can create different representations.

Define what is a car? It can drive, brake and steer.

```go
package car

type Car interface {
  Drive() error
  Brake() error
  Steer() error
}
```

What are the parts that make up a car?

```go
type Engine struct {
  PeakTorqueRPM uint
  MaxTorque     uint
}

type Wheels struct {
  Width    float64
  TireMake string
}
```

Now how should a builder behave? It should install parts.

```go
type Builder interface {
  Init(model string) Builder
  InstallEngine(*Engine) Builder
  InstallWheels(*Wheels) Builder
  Build() Car
}
```

Here's an example of builder that should satisfy the `Builder` interface. Let's say that Porsche can only build Boxster at the moment.

```go
type PorscheBuilder struct {
  vehicle *Boxster
}

func (pb *PorscheBuilder) Init(model string) {
  switch model {
  case Cayman:
      pb.vehicle = &Caymen{}
  case Boxster:
      pb.vehicle = &Boxster{}
  default:
      pb.vehicle = &Macan{}
  }
}

func (pb *PorscheBuilder) InstallEngine(e *Engine) Builder {
  pb.vehicle.horspower = float64(e.MaxTorque*e.PeakTorqueRPM) / 5252

  return pb
}

func (pb *PorscheBuilder) InstallWheels(w *Wheels) Builder {
  switch w.TireMake {
  case Michelin:
    pb.car.brakeDist = 16 * w.Width
  default:
    pb.car.brakeDist = 25 * w.Width
  }

  return pb
}

func (pb *PorscheBuilder) Build() Car {
  return pb.vehicle
}
```

The `Boxster` struct should look like this.

```go
type Boxster struct {
  horsepower float64
  brakeDist  float64
}

func (b *Boxster) Drive() error {
  fmt.Printf("Accelerating with %f horsepower\n", b.horsepower)
  return nil
}

func (b *Boxster) Brake() error {
  fmt.Printf("Coming to a stop in %f feet\n", b.brakeDist)
  return nil
}
```

Ideally we can have many different types of builder, like `BMWBuilder`, `AudiBuilder` and etc...

```go
func NewBuilder(brand string) (Builder, error) {
  switch brand {
  case Porsche:
    return &PorscheBuilder{}, nil
  case BMW:
    return nil, errors.New("can't build BMW yet")
  case Mazda:
    return nil, errors.New("can't build Mazda yet")
  default:
    return nil, fmt.Errorf("%s is not a recognized brand", brand)
  }
}
```

Finally we can use what we have developed to build some Boxsters.

```go
func main() {
  builder, err := car.NewBuilder(car.Porsche)
  if err != nil {
    return
  }

  car := builder.
    InstallWheels(&car.Wheels{Width: 7.5, TireMake: Michelin}).
    InstallEngine(&car.engine{MaxTorque: 300, PeakTorqueRPM: 6000}).
    Build()

  car.Drive()
  car.Brake()
}
```

