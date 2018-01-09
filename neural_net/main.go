package main

import (
	"fmt"
	"go-academy/neural_net/layer"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

func RandomMat(row, col int) *mat.Dense {
	randFloats := []float64{}
	for i := 0; i < row*col; i++ {
		randFloats = append(randFloats, rand.Float64())
	}

	return mat.NewDense(row, col, randFloats)
}

func main() {
	W1 := mat.NewDense(3, 3, []float64{1, 0, 0, 0, 1, 0, 0, 0, 1})
	B1 := mat.NewDense(3, 3, []float64{1, 0, 0, 1, 0, 0, 1, 0, 0})
	X := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	layer1 := layer.Affine{
		Weight: W1,
		Bias:   B1,
	}

	if result, err := layer1.ForwardProp(X); err == nil {
		fmt.Println(result)
	}

	gradOut := mat.NewDense(3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1})
	if gradX, gradW, gradB, err := layer1.BackwardProp(gradOut); err == nil {
		fmt.Println(gradX)
		fmt.Println(gradW)
		fmt.Println(gradB)
	}
}
