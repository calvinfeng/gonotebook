# Slice

## Add/Remove from Middle

Remove an element from the middle of a slice.

```golang
i := 1
s := append(s[:i], s[i+1:]...)
```

### Performance Consideration

The builtin `append()` needs to create a new backing array if the capacity of the destination slice
is less than what the length of the slice would be after the append. This also requires to copy the
current elements from destination to the newly allocated array, so there are much overhead.

## Sort

I can use Golang's sort package but my data type must implement the sort interface.

```golang
type Edge struct {
    Nodes  [2]*Node
    Weight int
}

// ByWeight implements Golang sort interface.
type ByWeight []*Edge

func (s ByWeight) Len() int {
    return len(s)
}

func (s ByWeight) Less(i, j int) bool {
    return s[i].Weight < s[j].Weight
}

func (s ByWeight) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
```
