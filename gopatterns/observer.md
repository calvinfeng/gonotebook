# Observer
The observer pattern allows an observable to emit events to observers. The emission can either be
done through channels or invoking registered callbacks.

```golang
package main

import "fmt"

type (
	Event struct {
		Data string
	}

	Observer interface {
		OnNotify(Event)
	}

	Observable interface {
		Register(Observer)
		Unregister(Observer)
		Notify(Event)
	}
)

type Reporter struct {
	observers map[Observer]struct{}
}

func (r *Reporter) Register(ob Observer) {
	r.observers[ob] = struct{}{}
}

func (r *Reporter) Unregister(ob Observer) {
	delete(r.observers, ob)
}

func (r *Reporter) Notify(e Event) {
	for ob := range r.observers {
		ob.OnNotify(e)
	}
}

type Audience struct {
	name string
}

func (a *Audience) OnNotify(e Event) {
	fmt.Printf("%s received breaking news! %s\n", a.name, e.Data)
}
```