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
	// GradOut := mat.NewDense(3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1})
	GradOut := RandomMat(3, 3)

	ForwardPropX := func(input *mat.Dense) (*mat.Dense, error) {
		aff := layer.Affine{
			Weight: W1,
			Bias:   B1,
		}

		return aff.ForwardProp(input)
	}

	ForwardPropW := func(weight *mat.Dense) (*mat.Dense, error) {
		aff := layer.Affine{
			Weight: weight,
			Bias:   B1,
		}

		return aff.ForwardProp(X)
	}

	ForwardPropB := func(bias *mat.Dense) (*mat.Dense, error) {
		aff := layer.Affine{
			Weight: W1,
			Bias:   bias,
		}

		return aff.ForwardProp(X)
	}
	fmt.Println("############################################Numerical Gradient####################################")

	if numGradX, err := layer.EvalNumericalGradient(ForwardPropX, X, GradOut, 1e-5); err == nil {
		fmt.Println(numGradX)
	}

	if numGradW, err := layer.EvalNumericalGradient(ForwardPropW, W1, GradOut, 1e-5); err == nil {
		fmt.Println(numGradW)
	}

	if numGradW, err := layer.EvalNumericalGradient(ForwardPropB, B1, GradOut, 1e-5); err == nil {
		fmt.Println(numGradW)
	}

	layer1 := layer.Affine{
		Weight: W1,
		Bias:   B1,
	}

	if result, err := layer1.ForwardProp(X); err == nil {
		fmt.Println("#######################################Forward Prop###########################################")
		fmt.Println(result)
	}

	if gradX, gradW, gradB, err := layer1.BackwardProp(GradOut); err == nil {
		fmt.Println("#########################################Back Prop############################################")
		fmt.Println(gradX)
		fmt.Println(gradW)
		fmt.Println(gradB)
	}
}
