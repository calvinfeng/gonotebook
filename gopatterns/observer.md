# Observer

The observer pattern allows an observable to emit events to observers. The emission can either be
done through channels or invoking registered callbacks.

## Conventional Approach

First let's define what are `Observer` and `Observable` as interfaces.

```go
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
```

Now we can create a `Reporter` like a news reporter that broadcast news to people.

```go
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
```

People are acting as `Audience` to the event that reporter reports.

```go
type Audience struct {
  name string
}

func (a *Audience) OnNotify(e Event) {
  fmt.Printf("%s received breaking news! %s\n", a.name, e.Data)
}
```

## Concurrent Approach

The problem with above aproach is that it is not concurrent. The observable calls `OnNotify` on each
observer. What if `OnNotify` is an expensive function? Then there will be a significant delay to
relay each event. We want to make it concurrent with room for parallelism.

```go
type Observable interface {
  Register(interface{})
  Unregister(interface{})
  Notify(data []byte)
}
```

Now we can register variable of any data type to this observable. We implement the interface like
before with few minor modifications. Notice that `Notify` is no longer calling a callback, it is
simply passing data to observers through channel. It is up to the observers to handle the data.

```go
func NewReporter() Observable {
  return &Reporter{
  observers: make(map[interface{}]chan []byte),
  }
}

type Reporter struct {
  observers map[interface{}]chan []byte
}

func (r *Reporter) Register(ob interface{}) chan []byte {
  r.observers[ob] = make(chan []byte)
  return r.observers[ob]
}

func (r *Reporter) Unregister(ob interface{}) {
  delete(r.observers, ob)
}

func (r *Reporter) Notify(data []byte) {
  for _, ch := range r.observers {
    select {
    case ch <- data:
    default:
    }
  }
}
```

Finally let's put it to use.

```go
type Person struct {
  Name string
}

func main() {
  ob := NewReporter() // This is our observable
  p := &Person{Name: "Calvin"} // This is our observer

  ch := ob.Register(p)
  go func() {
    for data := range ch {
      fmt.Println(p.Name, "received data", string(data))
    }
  }()

  timer := time.NewTimer(10 * time.Second)
  ticker := time.NewTicker(200 * time.Millisecond)
  for {
    select {
    case <-timer.C:
      return
    case <-ticker.C:
      ob.Notify([]byte("hello there"))
    }
  }
}
```