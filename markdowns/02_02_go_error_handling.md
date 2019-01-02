# Go Error Handling

Error handling is a must for any program. The common approach in JavaScript, Python and Java is
using try catch statement, but we don't have that in Go! For example, let's say we are trying to
feed Eric Cartman with unhealthy food.

In Python, we would do the following.

```python
def feed(kid, food):
  if food.calorie > 1000:
    raise ValueError("this will kill your kid")
  
  kid.eat(food)
  

def main():
  # Initialize the objects...
  try:
    feed(eric, french_fries)
  except ValueError:
    print "please don't do that, he's fat, not big bone"
```

In Go, we would simply return an error.

```golang
func feed(k *Kid, f *Food) error {
  if food.calorie > 1000 {
    return errors.New("this will kill your kid")
  }

  k.eat(f)
  return nil
}

func main() {
  // Initialize the objects...
  if err := feed(eric, frenchFries); err != nil {
    fmt.Println("please don't do that, he's fat, not big bone")
    log.Error(err) // Log the error
  }
}
```