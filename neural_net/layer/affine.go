package layer

import (
	"errors"
	"gonum.org/v1/gonum/mat"
)

type Affine struct {
	Input  *mat.Dense // Dense matrix of shape (N, D) where D is the feature dimension
	Weight *mat.Dense // Dense matrix of shape (D, H) where H is the hidden dimension
	Bias   *mat.Dense // Dense matrix of shape (1, H)
	Output *mat.Dense // Dense matrix of shape (N, H)
}

// Expected input is of shape (N, D) where N is the number of batch example
func (a *Affine) ForwardProp(X *mat.Dense) (*mat.Dense, error) {
	// First check dimension, expected output is N by H
	N, xCol := X.Dims()
	wRow, H := a.Weight.Dims()

	if xCol != wRow {
		return nil, mat.ErrShape
	}

	// Cache the input, cause we need it to compute back propagation
	a.Input = mat.NewDense(N, xCol, nil)
	a.Input.Copy(X)

	// Initialize a nil matrix
	result := mat.NewDense(N, H, nil)

	// Perform operations
	result.Mul(X, a.Weight)
	result.Add(result, MatBroadcast(a.Bias, N))

	return result, nil
}

// BackwardProp takes an upstream gradient, a.k.a. gradient of output and perform gradient calculation on layer's input,
// weight and bias matrices. The expected input to this function should be of shape, (N, H) which is the same shape as
// the output from ForwardProp.
func (a *Affine) BackwardProp(GradOut *mat.Dense) (*mat.Dense, *mat.Dense, *mat.Dense, error) {
	if a.Input == nil {
		return nil, nil, nil, errors.New("input is not set")
	}

	N, H := GradOut.Dims()
	xRow, xCol := a.Input.Dims()
	wRow, wCol := a.Weight.Dims()

	if N != xRow || H != wCol {
		return nil, nil, nil, mat.ErrShape
	}

	gradW := mat.NewDense(xCol, H, nil)
	gradW.Mul(a.Input.T(), GradOut)

	gradX := mat.NewDense(N, wRow, nil)
	gradX.Mul(GradOut, a.Weight.T())

	// Sum along column on upstream gradient
	gradB := mat.NewDense(1, H, SumAlongColumn(GradOut))

	return gradX, gradW, gradB, nil
}
