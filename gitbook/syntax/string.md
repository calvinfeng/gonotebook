# String

> In Go, a string is a read-only slice of bytes. 

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
