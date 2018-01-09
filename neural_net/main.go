package main

import (
	"fmt"
	"go-academy/neural_net/layer"
	"gonum.org/v1/gonum/mat"
)

func main() {
	W1 := mat.NewDense(3, 3, []float64{2, 0, 0, 0, 2, 0, 0, 0, 2})
	X := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	layer1 := layer.Affine{
		Weight: W1,
	}

	if result, err := layer1.ForwardProp(X); err == nil {
		fmt.Println(result)
	}
}
