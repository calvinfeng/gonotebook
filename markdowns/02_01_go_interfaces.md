# Go Interfaces
Interface is a very powerful tool for creating abstractions in Go. It also enables us to create mock
for unit testing. Since Go is strongly typed, every function expects a concrete type. For example,
```golang
func Feed(d *Dog) bool {
    // Return true if the dog eats it else false.
}
```

The `Feed` function can only feed dogs due to the strong typed nature of the language. One may naturally
ask, *what if I want to feed a cat?*
```golang
func FeedCat(c *Cat) bool {
    // Implementation...
}

func FeedDog(d *Dog) bool {
    // Implementation...
}
```

That obviously does not look good because in practice we would like to write one function that feeds
as many animals as possible. So `interface` comes to rescue. We define an animal interface.
```golang
type Food struct{}

type Animal interface{
    Eat(Food)
}
```

Any data structure that implements the function `Eat` with argument `Food` is said to be satisfying
the `Animal` interface.
```golang
type Dog struct{}

func (d *Dog) Eat(f Food) {
    // Eat...
}

type Cat struct{}

func (c *Cat) Eat(f Food) {
    // Eat
}
```

Now we just need to modify the `Feed` function a little bit.
```golang
func Feed(a Animal) bool {
    // Now we can feed anything that has Eat()
}
```

## Unit Tests
Suppose we want to isolate the `Feed` function unit tests from the `Animal` unit tests, we can mock
the animals!
```golang
type MockDog struct{}

func (md *MockDog) Eat(f Food) {
    assert.True(f.Volume > 100)
    asssert.Called()
}
```

And then pass it into `Feed` and see if test passes.
```
md := &MockDog
Feed(md)
``` 