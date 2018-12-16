package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	leftOp, leftErr := strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
	rightOp, rightErr := strconv.ParseFloat(r.URL.Query().Get("rop"), 64)

	result := fmt.Sprintf("%v + %v = %v", leftOp, rightOp, leftOp+rightOp)
	if leftErr == nil && rightErr == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid query parameters"))
	}
}

func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	leftOp, leftErr := strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
	rightOp, rightErr := strconv.ParseFloat(r.URL.Query().Get("rop"), 64)

	result := fmt.Sprintf("%v - %v = %v", leftOp, rightOp, leftOp-rightOp)
	if leftErr == nil && rightErr == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid query parameters"))
	}
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
