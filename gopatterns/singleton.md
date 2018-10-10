# Singleton
This one is pretty common and clear.

> Singleton creational design pattern restricts the instantiation of a type to a single object.

Previously I built an army, now I wish to have a general that will lead my army. I just want one 
general because I don't want a mutiny.
```golang 
package army

var (
    once sync.Once
    commander *General
)

// Outside of this package, no one can access this struct except through NewGeneral()` 
type general struct {
    Army Army
}

func NewGeneral(armySize int) *general {
    once.Do(func() {
        commander = &general{Army: New(armySize)}
    })

    return commander
}
```