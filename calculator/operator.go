package main

import "fmt"

// Operator is a function that implements arithmetic operator. The return value is a string that
// expresses the operation, like 1 + 1 = 2.
type Operator func(lop float64, rop float64) string

// Addition is adding two numbers together.
func Addition(lop float64, rop float64) string {
	return fmt.Sprintf("%v + %v = %v", lop, rop, lop+rop)
}

// Subtraction is subtracting right operand from left operand.
func Subtraction(lop float64, rop float64) string {
	return fmt.Sprintf("%v - %v = %v", lop, rop, lop-rop)
}

// Multiplication is multiplying two numbers together.
func Multiplication(lop float64, rop float64) string {
	return fmt.Sprintf("%v * %v = %v", lop, rop, lop*rop)
}

// Division is divding left operand by right operand.
func Division(lop float64, rop float64) string {
	return fmt.Sprintf("%v / %v = %v", lop, rop, lop/rop)
}
