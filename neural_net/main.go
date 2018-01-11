package main

import (
	"fmt"
	"go-academy/neural_net/layer"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func main() {
	// If seed is not set, we will expect to get the result random numbers in runtime every time we run the program. This
	// particular seed will give the program pseudo-randomness.
	rand.Seed(time.Now().UTC().Unix())

	W1 := mat.NewDense(3, 3, []float64{1, 0, 0, 0, 1, 0, 0, 0, 1})
	B1 := mat.NewDense(3, 3, []float64{1, 0, 0, 1, 0, 0, 1, 0, 0})
	X := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// GradOut := mat.NewDense(3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1})
	GradOut := RandNormMat(3, 3, 1e-2, 0)

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

	if numGradX, err := layer.EvalNumericalGrad(ForwardPropX, X, GradOut, 1e-5); err == nil {
		fmt.Println(numGradX)
	}

	if numGradW, err := layer.EvalNumericalGrad(ForwardPropW, W1, GradOut, 1e-5); err == nil {
		fmt.Println(numGradW)
	}

	if numGradB, err := layer.EvalNumericalGradForBias(ForwardPropB, B1, GradOut, 1e-5); err == nil {
		fmt.Println(numGradB)
	}

	layer1 := layer.Affine{
		Weight: W1,
		Bias:   B1,
	}

	if _, forwardPropErr := layer1.ForwardProp(X); forwardPropErr == nil {
		if gradX, gradW, gradB, backPropErr := layer1.BackwardProp(GradOut); backPropErr == nil {
			fmt.Println(gradX)
			fmt.Println(gradW)
			fmt.Println(gradB)
		}
	}

	layer2 := layer.Sigmoid{}

	if out, forwardPropErr := layer2.ForwardProp(X); forwardPropErr == nil {
		fmt.Println(out)
	}
}
