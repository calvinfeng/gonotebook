package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Operator is a function that performs arithmetic operation.
type Operator func(lop float64, rop float64) string

// Addition adds two numbers together.
func Addition(lop float64, rop float64) string {
	return fmt.Sprintf("%v + %v = %v", lop, rop, lop+rop)
}

// Subtraction subtracts one number from another.
func Subtraction(lop float64, rop float64) string {
	return fmt.Sprintf("%v - %v = %v", lop, rop, lop-rop)
}

// Multiplication multiples two numbers.
func Multiplication(lop float64, rop float64) string {
	return fmt.Sprintf("%v * %v = %v", lop, rop, lop*rop)
}

// Division divides one number from another.
func Division(lop float64, rop float64) string {
	return fmt.Sprintf("%v / %v = %v", lop, rop, lop/rop)
}

// HandleAdd is a handler that adds two query params together.
func HandleAdd(w http.ResponseWriter, r *http.Request) {
	leftOp, leftErr := strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
	rightOp, rightErr := strconv.ParseFloat(r.URL.Query().Get("rop"), 64)

	result := fmt.Sprintf("%v + %v = %v", leftOp, rightOp, leftOp+rightOp)
	if leftErr != nil || rightErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid query parameters"))
		return

	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

// HandleSubtract is a handler that subtracts one query param from another.
func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	leftOp, leftErr := strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
	rightOp, rightErr := strconv.ParseFloat(r.URL.Query().Get("rop"), 64)

	result := fmt.Sprintf("%v - %v = %v", leftOp, rightOp, leftOp-rightOp)
	if leftErr != nil || rightErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid query parameters"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

// CreateOperationHandler dynamically creates new handler for different type of operations.
func CreateOperationHandler(op Operator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var lop, rop float64
		var err error

		lop, err = strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid lop query parameter"))
			return
		}

		rop, err = strconv.ParseFloat(r.URL.Query().Get("rop"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid rop query parameter"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(op(lop, rop)))
	}
}
