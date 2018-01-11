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
	X := RandNormMat(3, 3, 1, 0)
	W1 := RandNormMat(3, 3, 1, 0)
	B1 := RandNormMat(3, 3, 1e-2, 0)
	GradOut := RandNormMat(3, 3, 1e-2, 0)

	AffineForwardPropX := func(input *mat.Dense) (*mat.Dense, error) {
		aff := layer.Affine{
			Weight: W1,
			Bias:   B1,
		}

		return aff.ForwardProp(input)
	}

	SigForwardProp := func(input *mat.Dense) (*mat.Dense, error) {
		sig := layer.Sigmoid{}

		return sig.ForwardProp(input)
	}

	AffineForwardPropW := func(weight *mat.Dense) (*mat.Dense, error) {
		aff := layer.Affine{
			Weight: weight,
			Bias:   B1,
		}

		return aff.ForwardProp(X)
	}

	AffineForwardPropB := func(bias *mat.Dense) (*mat.Dense, error) {
		aff := layer.Affine{
			Weight: W1,
			Bias:   bias,
		}

		return aff.ForwardProp(X)
	}

	if numGradX, err := layer.EvalNumericalGrad(AffineForwardPropX, X, GradOut, 1e-5); err == nil {
		fmt.Println(numGradX)
	}

	if numGradW, err := layer.EvalNumericalGrad(AffineForwardPropW, W1, GradOut, 1e-5); err == nil {
		fmt.Println(numGradW)
	}

	if numGradB, err := layer.EvalNumericalGradForBias(AffineForwardPropB, B1, GradOut, 1e-5); err == nil {
		fmt.Println(numGradB)
	}

	if numGradX, err := layer.EvalNumericalGrad(SigForwardProp, X, GradOut, 1e-5); err == nil {
		fmt.Println(numGradX)
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

	if _, forwardPropErr := layer2.ForwardProp(X); forwardPropErr == nil {
		if gradX, backPropErr := layer2.BackwardProp(GradOut); backPropErr == nil {
			fmt.Println(gradX)
		}
	}
}
