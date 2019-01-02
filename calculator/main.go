package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/api/add", CreateOperationHandler(Addition))
	http.HandleFunc("/api/sub", CreateOperationHandler(Subtraction))
	http.HandleFunc("/api/mul", CreateOperationHandler(Multiplication))
	http.HandleFunc("/api/div", CreateOperationHandler(Division))
	fmt.Println("starting calculator server on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("failed to serve", err)
	}
}

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

type Operator func(lop float64, rop float64) string

func Addition(lop float64, rop float64) string {
	return fmt.Sprintf("%v + %v = %v", lop, rop, lop+rop)
}

func Subtraction(lop float64, rop float64) string {
	return fmt.Sprintf("%v - %v = %v", lop, rop, lop-rop)
}

func Multiplication(lop float64, rop float64) string {
	return fmt.Sprintf("%v * %v = %v", lop, rop, lop*rop)
}

func Division(lop float64, rop float64) string {
	return fmt.Sprintf("%v / %v = %v", lop, rop, lop/rop)
}

func CreateOperationHandler(op Operator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var lop, rop float64
		var err error

		lop, err = strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid left operand"))
			return
		}

		rop, err = strconv.ParseFloat(r.URL.Query().Get("rop"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid right operand"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(op(lop, rop)))
	}
}
