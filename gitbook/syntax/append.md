# Golang Append

## Arrays

The size of an array is part of its type. This limits the expressive power a bit. However, sometimes
limiting expression is what one wants, just like a type system.

```go
var buffer [256]byte
```

The type of buffer includes its size, `[256]byte`. An array with 512 bytes would be of the distinct
type `[512]byte`.

Arrays have their place, they are a good representation of a transformation matrix for instance, but
their most common purpose in Go is to hold storage for a slice.

## Slices

A _slice_ is a data structure describing a contiguous section of an array stored separately from the
slice variable itself. 

> A slice is not an array. A slice describes a piece of an array.

We can **slice** the buffer to create a slice.

```go
var slice []byte = buffer[100:150]
```

What is a slice? It's not quite the full story but think of it as a data structure with two
elements, a length and a pointer to an element of an array.

```go
type sliceHeader struct {
    Length: int
    ZerothElement *byte
}

slice := sliceHeader{
    Length:        50,
    ZerothElement: &buffer[100],
}
```

The slice header is not visible to the programmer, and the type of the element pointer depends on
the type of the elements, but this gives the general idea of the mechanics.

We can slice a slice.

```go
slice2 := slice[5:10] 
```

This operation creates a new slice, in this case with elements 5 through 9 of the original slice, 
which means elements 105 through 109 of the original array. The underlying slice header would look
like this.

```go
slice2 := sliceHeader{
    Length:        5,
    ZerothElement: &buffer[105],
}
```

The header is still pointing to the same underlying array, stored in the buffer variable.

## Slice Passes as Value

Although slice contains a pointer, it is itself a value. Under the covers, it is a struct value
holding a pointer and a length. It is **NOT** a pointer to a struct.

Passing a slice to a function is essentially copying a slice header. I am using slice and slice
header interchangeably because it's the language the original Go post uses.

```go
func reduceLength(slice []byte) []byte {
    slice = slice[0:len(slice)-1]
    return slice
}
```

## Slice Passes as Reference

Similar to previous example, instead of returning a new slice, we can modify the original slice if
we accept a pointer.

```go
func reduceLength(slicePtr *[]byte) {}
    slice := *slicePtr
    *slicePtr = slice[0:len(slice)-1]
}
```

Let's say we wanted to have a method on a slice that truncates it as the final slash.

```go
type path []byte

func (p *path) {
    i := bytes.LastIndex(*p, []byte("/"))
    if i >= 0 {
        *p = (*p)[0:i]
    }
}
```

## Capacity

Besides the array pointer and length, the slice header also stores its capacity. The capacity field
records how much space the underlying array actually has; it is the maximum value the length can
reach. Trying to grow the slice beyond its capacity will step beyond the limits of the array and
will trigger a panic.

```go
slice := buffer[0:0]
```

The header looks like this.

```go
slice := sliceHeader{
    Length: 0,
    Capacity: 10,
    ZerothElement: &buffer[0],
}
```

## Make

If we want to grow a slice beyond its capacity, we need to define a new underlying array and copy
over the data. The built-in function `make` allocates a new array and creates a slice header to
describe it.

## Copy

Go has a built-in function to copy arrays. The copy function is smart. It only copies what it can,
paying attention to the lengths of both arguments. In other words, the number of elements it copies
is the minimum length of the two slices.

Assuming a slice has enough room for new element, an insert function would write like this.

```go
func insert(slice []int, index, value int) []int {
    slice = slice[0:len(slice)+1]
    copy(slice[index+1:], slice[index:])
    slice[index] = value
    return slice
}
```

## Append

We have enough ingredients to build the `append` function.

```go
func extend(slice []int, element int) []int {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]int, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:n+1]
    slice[n] = element
    return slice
}

func append(slice []int, items ...int) []int {
    for _, item := range items {
        slice = extend(slice, item)
    }
    return slice
}
```

We can make the `append` more efficient by allocating no more than once.

```go
func append(slice []int, elements ...int) []int {
    n := len(slice)
    total := len(slice) + len(elements)
    if total > cap(slice) {
        // Reallocate. Grow to 1.5 times the new size, so we can still grow.
        newSize := total*3/2 + 1
        newSlice := make([]int, total, newSize)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[:total]
    // Values are copied into the underlying array, despite multiple re-slicing.
    copy(slice[n:], elements)
    return slice
}
```

## Built-in Append

Since slice header is always updated by a call to `append`, we need to save the returned slice
after the call. In fact, the compiler won't let us call append without saving the result!

All slice tricks can be found here https://github.com/golang/go/wiki/SliceTricks.

## Nil Slice

Nil slice is essentially the zero value of slice header.

```go
sliceHeader {
    Lenght: 0,
    Capacity: 0,
    ZerothElement: nil
}
```

## String

Strings are read-only slices of bytes with a bit of extra syntactic support from the language. They
are read-only, there is no need for a capacity because we can't grow them anyways.

The array underlying a string is hidden from view; there is no way to access its contents except
through the string. That means that when we do string to bytes conversion, a copy of the array must
be made. Modification of the array underlying the byte slice don't affect the corresponding string.

```
str := string(slice)
slice := []byte(usr)
```
