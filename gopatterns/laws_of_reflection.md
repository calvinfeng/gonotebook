# The Laws of Reflection

> Reflection in computing is the ability of a program to examine its own structure, particularly
> through types; it's a form of metaprogramming.

## Interface

Reflection canot be explained without a good grasp on what is an interface in Go. I am not going to
show an example of how to use interface because this should be quite trivial. A variable of interface
type stores two key pieces of information.

1. Concrete value that is assigned to the variable.
2. Concrete type of the value.

We can extract the value and type of an interface variable using `reflect.ValueOf` adn `reflection.TypeOf`.

For example,

```golang
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
}

func main() {
    var myVar interface{}

    myVar = 1

    fmt.Println(reflect.ValueOf(myVar)) // => 1
    fmt.Println(reflect.TypeOf(myVar)) // => int

    myVar = Person{"Calvin"}

    fmt.Println(reflect.ValueOf(myVar)) // => {Calvin}
    fmt.Println(reflect.TypeOf(myVar)) // => main.Person
}
```

## First Law

> Reflection goes from interface value to reflection object.

### Reflection Object(s)

There are two types of reflection objects, `reflect.Value` and `reflect.Type`.

1. `reflect.Value` - Representation of runtime data
2. `reflect.Type` - Representation of a Go type

### Kind

Reflection objects have a method called `Kind()` which allows us to examine what sort of items is
stored in the variable.

Here's a list of possible `Kind` constants.

```golang
type Kind uint

const (
    Invalid Kind = iota
    Bool
    Int
    Int8
    Int16
    Int32
    Int64
    Uint
    Uint8
    Uint16
    Uint32
    Uint64
    Uintptr
    Float32
    Float64
    Complex64
    Complex128
    Array
    Chan
    Func
    Interface
    Map
    Ptr
    Slice
    String
    Struct
    UnsafePointer
)
```

We can easily check if a variable is a pointer, slice, string or anything by equal comparison. Also
it is not necessary to convert our variable into `interface{}` type before we do reflection because
the method `reflect.ValueOf` and `reflect.TypeOf` will do the work for us when we pass in the argument.

For example,

```golang
num := 123
fmt.Println(reflect.ValueOf(&num).Kind()) // => reflect.Ptr
fmt.Println(reflect.TypeOf(&num).Kind()) // => reflect.Ptr
```

### Elem

If we have a pointer value, we may want to extract the actual value of the pointer. We can use `Elem`
in this case.

For example,

```golang
num := 123
fmt.Println(reflect.ValueOf(&num)) // => 0xc00001a110
fmt.Println(reflect.ValueOf(&num).Elem()) // => 123
fmt.Println(reflect.TypeOf(&num)) // => *int
fmt.Println(reflect.TypeOf(&num).Elem()) // => int
```

However, be aware that `Elem()` of a non-pointer will cause panic.

### Explicit Value

We can grab the explicit value of a `reflect.Value` by using the following list of methods.

1. `String()`
2. `Slice()`
3. `Uint()`
4. `Int()`
5. `Float()`
6. `Pointer()`
7. `Complex()`
8. `Bool()`
9. The list goes on, there may be other data types I am not listing here.

For example,

```golang
num := 123
fmt.Println(reflect.ValueOf(num).Int()) // => 123
```

### NumFields

If we have a `reflect.Type` that is describing the concrete type of a struct variable, we can even
examine the number of fields this struct has.

```golang
type Contact struct {
    Name    string
    Phone   int
    Address string
}

func main() {
    c := Contact{"Calvin", 4151231234, "123 Example Street"}
    for i := 0; i < reflect.TypeOf(c).NumField(); i++ {
        fmt.Println(reflect.TypeOf(c).Field(i).Name) // => Name, Phone, Address
        fmt.Println(reflect.ValueOf(c).Field(i)) // => Calvin, 4151231234, 123 Example Street
    }
}
```

## Second Law

> Reflection goes from reflection object to interface value.

Given a `reflect.Value`, we can recover the interface using `Interface()`. The method packs the value
and type information back into an interface representation and returns the result.

For example,

```golang
func main() {
    origin := &Contact{"Calvin", 4151231234, "123 Example Street"}
    c, ok := reflect.ValueOf(origin).Interface().(*Contact)
    if ok {
        fmt.Println(c) // => &{Calvin, 4151231234, 123 Example Street}
    }
}
```

In other words, `Interface()` is the inverse of `ValueOf()`, except that the returned value is always
of an empty interface type. We will need to perform type assertion to restore the original type. if
we intend to use the variable in subsequeuent logic.

## Third Law

> To modify a reflection object, the value must be settable.

Settability is a bit like addressability, but stricter. It's the property that a reflection object
can modify the actual storage that was used to create the reflection object. Settability is determined
by whether the reflection object holds the original item.

Let's look at an example.

```golang
x := 10
reflect.ValueOf(x).SetInt(11)
```

The code above will result in,

    panic: reflect: reflect.Value.SetInt using unaddressable value.

The variable `x` is *unsettable*.

```golang
reflect.ValueOf(x).CanSet() // => false
```

It is because we are passing a copy of x into ValueOf. Thus, if the set statement were allowed to
succeed, it would not update x even though v looks like it was created from x. The Golang creator
decided that this would be confusing and declared it as illegal.

The best way to think about this is to think of us passing `x` into a function.

```golang
x := 10

func Increment(x int) {
    x += 1
}

Increment(x)
fmt.Println(x == 11) // => false
```

If we want it to be settable, we must pass in a pointer or a reference. However, the following does
not work.

```golang
reflect.ValueOf(&x).SetInt(11)
```

Pointer to an `int` is still an integer. We need to get the element of the pointer to perform a
successful set operation. This is because we are not trying to set pointer to some value. We are
trying to set the value of the pointer to some value.

```golang
reflect.ValueOf(&x).Elem().SetInt(11)
```

We can now extend the example to a struct.

```golang
func main() {
    c := &Contact{"Calvin", 4151231234, "123 Example Street"}

    rv, rt := reflect.ValueOf(c), reflect.TypeOf(c)
    for i := 0; i < rt.Elem().NumField(); i++ {
        fmt.Printf("modifying %v\n", rt.Elem().Field(i).Name)
        if rv.Elem().Field(i).CanSet() && rv.Elem().Field(i).Kind() == reflect.String {
            rv.Elem().Field(i).SetString("foo")
        }
    }
}
```

The `Contact` struct will become,

    &{foo 4151231234 foo}
