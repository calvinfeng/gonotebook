# Bit Operators in Go

```
& bitwise AND
| bitwise OR
^ bitwise XOR
&^ AND NOT
<< left shift
>> right shift
```

For the purpose of demonstration, define an unsigned 8-bits type as the bit array or a single byte.

```go
type bits uint8
```

```go
var small bits = 1 // This is okay
var overflow bits = 256 // This will overflow
```

## Bitwise AND

```
AND(a, b) = 1 if a = b = 1 else 0
```

For example, 

```go
var a bits = 10
var b bit = 12

a & b // => 1010 & 1100 = 1000
```

## Bitwise OR

```
OR(a, b) = 1 if a = 1 or = 1 else 0
```

For example,

```go
var a bits = 10
var b bit = 12

a | b // => 1010 | 1100 = 1110
```

## Bitwise XOR (Exclusive OR)

```
XOR(a, b) = 1 if a != b else 0
```

For example,

```go
var a bits = 10
var b bit = 12

a ^ b // => 1010 ^ 1100 = 0110
```

## Bitwise AND NOT

```
AND_NOT(a, b) = AND(a, NOT(b))
```

It has an interesting property of clearing bits if second operand is 1. It does no operation if the
second operand is 0.

```go
var a bits = 10
var b bit = 12

a &^ b // => 1010 &^ 1100 = 0010

var a bits = 12
var b bits = 255

a &^ b // => 00001100 &^ 11111111 = 0
```

## Shift

- `<<` Shift left means shifting all bits to the left by N times. 
- `>>` Shift right means shifting all bits to the right by N times.

For example,

```go
var a = 10

a << 1 // => 1010 << 1 = 10100
a << 2 // => 1010 << 2 = 101000
a >> 1 // => 1010 >> 1 = 101
a >> 2 // => 1010 >> 2 = 10
```
