# Object Pool

The object pool creational design pattern is used to prepare and keep multiple instances according
to the demand expectation.

## Advantage

Object pool pattern is useful in cases where object initialization is more expensive than object
maintenance. It has positive effects on performance due to objects being initialized beforehand.

## Disadvantage

If there are spikes in demand as opposed to a steady demand, the maintenance overhead might outweigh
the benefits.

## Example

For example, I have solider which requires *training* before I can initialize it.

```go
package army

// NewSoldier trains and creates a soldier.
func NewSoldier(id uint) *Soldier {
  for i := 1; i <= 100; i++ {
    fmt.Printf("\rTraining %d/100", i)
    time.Sleep(time.Millisecond)
  }
  fmt.Printf("\nDone\n")

  return &Soldier{ID: id, Skill: 100}
}

type Soldier struct {
  ID    uint
  Skill uint
}
```

I can create an army using a channel. Treat this as a thread-safe queue.

```go
type Army chan *Soldier

func New(total int) Army {
  a := make(Army, total)
  for i := 0; i < total; i++ {
    a <- NewSoldier(uint(i + 1))
  }

  return a
}
```

Then I can use the army in my main function.

```go
package main

func main() {
    a := army.New(100)

    // Summon soldier in parallel
    for i := 0; i < 100; i++ {
        go func() {
            s := <-a
            s.March()
            a <- s
        }()
    }
}
```