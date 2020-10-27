# String

> In Go, a string is a read-only slice of bytes. 

## []byte vs []rune

It's peculiar that indexing a string returns byte but ranging over a string, each element is rune.
Consider this example,

```go
str := "中文"
fmt.Println(len(str)) // => 6 not 2
```

I would expect 2 to be printed to the screen but it's actually 6. Each Chinese character is
represented by 3 bytes. Indexing into a string is indexing into each individual byte. However,
ranging over the string allows the bytes to be packed into one code point, a.k.a rune or int32.

The reason is simple, in order to for an indexing function to return a rune, it needs to scan the
string to allow multiple bytes to be packed together to form one character code point. This will
force indexing to be O(n) instead of O(1).

On the other hand, ranging over a string is already performing an O(n) operation so the language can
provide a bit convenience.

```go
str := "中文"
var count int
for range str {
    count++
}
fmt.Println(count) // => 2 not 6
```

## Enumeration

String has to be converted to `[]rune` for iteration. Golang uses UTF-8 encoding for all strings, even
the source code is in UTF-8. Each character is 4 bytes, which is represented by a 32-bit integer, 
known as `rune`. 32-bit integer can represent a lot, a lot, a lot of characters from different
languages.

```go
s = "hello world"
for pos, r := range s {
    fmt.Printf("character %c starts at byte position %d\n", char, pos)
}
```

## Append/Insert

Since strings are slices of bytes, the operation is similar to that of slices.

```go
s := "abc"
r := "def"
result := string(append([]rune(s), []rune(r)...))
```

Insert a new character in between.

```go
original := "abcef"
result := original[:3] + "d" + original[3:]
```

