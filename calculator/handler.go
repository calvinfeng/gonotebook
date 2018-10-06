package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	ADD = "ADD"
	SUB = "SUBTRACT"
	DIV = "DIVIDE"
	MUL = "MULTIPLY"
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

func OperationHandlerCreator(opType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		leftOp, leftErr := strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
		rightOp, rightErr := strconv.ParseFloat(r.URL.Query().Get("rop"), 64)

		if leftErr != nil || rightErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid query parameters."))
			return
		}

		var result string
		switch opType {
		case ADD:
			result = fmt.Sprintf("%v + %v = %v", leftOp, rightOp, leftOp+rightOp)
			break
		case SUB:
			result = fmt.Sprintf("%v - %v = %v", leftOp, rightOp, leftOp-rightOp)
			break
		case MUL:
			result = fmt.Sprintf("%v * %v = %v", leftOp, rightOp, leftOp*rightOp)
			break
		case DIV:
			result = fmt.Sprintf("%v / %v = %v", leftOp, rightOp, leftOp/rightOp)
			break
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unrecognized operation."))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}
