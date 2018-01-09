package layer

import (
	"gonum.org/v1/gonum/mat"
)

type Affine struct {
	Input  *mat.Dense
	Output *mat.Dense
	Weight *mat.Dense
}

// TODO: Add bias
func (a *Affine) ForwardProp(x *mat.Dense) (*mat.Dense, error) {
	// First check dimension
	xRow, xCol := x.Dims()
	wRow, wCol := a.Weight.Dims()

	if xCol != wRow {
		return nil, mat.ErrShape
	}

	// Cache the input, cause we need it to compute back propagation
	a.Input = mat.NewDense(xRow, xCol, nil)
	a.Input.Copy(x)

	result := mat.NewDense(xRow, wCol, nil)
	result.Mul(x, a.Weight)

	return result, nil
}

// TODO: Add Bias
func (a *Affine) BackwardProp(gradOut *mat.Dense) (*mat.Dense, *mat.Dense, error) {
	if a.Input == nil {
		return nil, nil, mat.Error{}
	}

	gradOutRow, gradOutCol := gradOut.Dims()
	xRow, xCol := a.Input.Dims()
	wRow, wCol := a.Weight.Dims()

	if gradOutRow != xRow || gradOutCol != wCol {
		return nil, nil, mat.ErrShape
	}

	gradW := mat.NewDense(xCol, gradOutCol, nil)
	gradW.Mul(a.Input.T(), gradOut)

	gradX := mat.NewDense(gradOutRow, wRow, nil)
	gradX.Mul(gradOut, a.Weight.T())

	return gradX, gradW, nil
}
